import radius as radius

@description('The URL of the server hosting test Terraform modules.')
param moduleServer string = 'http://localhost:8000'

@description('Name of the Redis Cache resource.')
param redisCacheName string = 'redis-rf-db'

@description('Name of the Radius Application.')
param appName string = 'tf-test-redis1'

resource env 'Applications.Core/environments@2023-10-01-preview' = {
  name: 'corerp-resources-terraform-redis-env2'
  properties: {
    compute: {
      kind: 'kubernetes'
      resourceId: 'self'
      namespace: 'corerp-resources-terraform-redis-env2'
    }
    recipes: {
      'Applications.Core/extenders': {
        default: {
          templateKind: 'terraform'
          templatePath: 'http://localhost:8000/kubernetes-redis.zip'
        }
      }
    }
  }
}

resource app 'Applications.Core/applications@2023-10-01-preview' = {
  name: appName
  properties: {
    environment: env.id
    extensions: [
      {
        kind: 'kubernetesNamespace'
        namespace: appName
      }
    ]
  }
}

resource webapp 'Applications.Core/extenders@2023-10-01-preview' = {
  name: 'corerp-resources-terraform-redis2'
  properties: {
    application: app.id
    environment: env.id
    recipe: {
      name: 'default'
      parameters: {
        redis_cache_name: redisCacheName
      }
    }
  }
}
