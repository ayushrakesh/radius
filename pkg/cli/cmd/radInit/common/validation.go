// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package common

import (
	"strings"

	"github.com/project-radius/radius/pkg/cli/prompt"
	corerp "github.com/project-radius/radius/pkg/corerp/api/v20220315privatepreview"
	"github.com/spf13/cobra"
)

const (
	SelectExistingEnvironmentPrompt         = "Select an existing environment or create a new one"
	SelectExistingEnvironmentCreateSentinel = "[create new]"
	EnterEnvironmentNamePrompt              = "Enter an environment name"
)

// SelectExistingEnvironment prompts the user to select from existing environments (with the option to create a new one).
// We also expect the the existing environments to be a non-empty list, callers should check that.
//
// If the name returned is empty, it means that that either no environment was found or that the user opted to create a new one.
func SelectExistingEnvironment(cmd *cobra.Command, defaultVal string, prompter prompt.Interface, existing []corerp.EnvironmentResource) (string, error) {
	// On this code path, we're going to prompt for input.
	//
	// Build the list of items in the following way:
	//
	// - default environment (if it exists)
	// - (all other existing environments)
	// - [create new]
	items := []string{}
	for _, env := range existing {
		if strings.EqualFold(defaultVal, *env.Name) {
			items = append(items, defaultVal)
			break
		}
	}
	for _, env := range existing {
		// The default is already in the list
		if !strings.EqualFold(defaultVal, *env.Name) {
			items = append(items, *env.Name)
		}
	}
	items = append(items, SelectExistingEnvironmentCreateSentinel)

	choice, err := prompter.GetListInput(items, SelectExistingEnvironmentPrompt)
	if err != nil {
		return "", err
	}

	if choice == SelectExistingEnvironmentCreateSentinel {
		// Returing empty tells the caller to create a new one.
		return "", nil
	}

	return choice, nil
}