# mvm-go-cloudsql
An example app for how to connect to Cloud SQL Second Generation on Managed VMs

See https://cloud.google.com/sql/docs/sql-proxy#gae for more information

This app connects to an instance named 'instance-name' located in a region
'region' and inside a project named 'project-id'. These represent the
placeholders and should be substituted with real values before the app is
deployed.

For example, for an instance named 'production' in 'us-central1' and in a project
named 'carrotmans-awesome-project', one would substitute occurrences of
`project-id:region:instance-name` with
`carrotmans-awesome-project:us-central1:production`. These references exist in two
places: in the app.yaml and in the definition of the Data Source Name (dsn) in
module1.go.

Or, an automated way. Note that you should still alter the INSTANCE, REGION, and
PROJECT variables to match your reality:

    INSTANCE="production"
    REGION="us-central1"
    PROJECT="carrotmans-awesome-project"
    
    git clone https://github.com/Carrotman42/mvm-go-cloudsql.git
    cd mvm-go-cloudsql
    sed -i -e "s/project-id:region:instance-name/$PROJECT:$REGION:$INSTANCE/" *
    gcloud --project=$PROJECT preview app deploy --no-promote app.yaml
