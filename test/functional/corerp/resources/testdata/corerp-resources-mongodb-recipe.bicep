import radius as radius

param rg string = resourceGroup().name

param sub string = subscription().subscriptionId

param magpieimage string 

resource env 'Applications.Core/environments@2022-03-15-privatepreview' = {
  name: 'corerp-resources-environment-recipes-env'
  location: 'global'
  properties: {
    compute: {
      kind: 'kubernetes'
      resourceId: 'self'
      namespace: 'corerp-resources-environment-recipes-env'
    }
    providers: {
      azure: {
        scope: '/subcriptions/${sub}/resourceGroup/${rg}'
      }
    }
    recipes: {
      mongodb: {
          connectorType: 'Applications.Connector/mongoDatabases' 
          templatePath: 'radiusdev.azurecr.io/recipes/mongodatabases/azure:1.0' 
      }
    }
  }
}

resource app 'Applications.Core/applications@2022-03-15-privatepreview' = {
  name: 'corerp-resources-mongodb-recipe'
  location: 'global'
  properties: {
    environment: env.id
  }
}

resource webapp 'Applications.Core/containers@2022-03-15-privatepreview' = {
  name: 'mongodb-recipe-app-ctnr'
  location: 'global'
  properties: {
    application: app.id
    connections: {
      mongodb: {
        source: recipedb.id
      }
    }
    container: {
      env: {
        DBCONNECTION: recipedb.connectionString()
      }
      image: magpieimage
    }
  }
}

resource recipedb 'Applications.Connector/mongoDatabases@2022-03-15-privatepreview' = {
  name: 'mongo-recipe-db'
  location: 'global'
  properties: {
    application: app.id
    environment: env.id
    recipe: {
      name: 'mongodb'
    }
  }
}