/*
Copyright 2023 The Radius Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package terraform

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	install "github.com/hashicorp/hc-install"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/hc-install/src"
	"github.com/project-radius/radius/pkg/ucp/ucplog"
)

const (
	installSubDir = "install"
)

// Install installs Terraform under /install in the provided Terraform root directory for the resource.
// Returns the path to the installed Terraform binary and the installer.
func Install(ctx context.Context, installer *install.Installer, tfDir string) (string, error) {
	logger := ucplog.FromContextOrDiscard(ctx)

	// Create Terraform installation directory
	installDir := filepath.Join(tfDir, installSubDir)
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory for terraform installation for resource: %w", err)
	}

	logger.Info(fmt.Sprintf("Installing Terraform in the directory: %q", installDir))
	// Re-visit this: consider checking if an existing installation of same version of Terraform is available.
	// For initial iteration we will always install Terraform for every execution of the recipe driver.
	execPath, err := installer.Ensure(ctx, []src.Source{
		&releases.LatestVersion{
			Product:    product.Terraform,
			InstallDir: installDir,
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to install terraform: %w", err)
	}
	logger.Info(fmt.Sprintf("Terraform latest version installed to: %q", execPath))

	return execPath, nil
}