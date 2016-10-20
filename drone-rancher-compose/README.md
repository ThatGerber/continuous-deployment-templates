Set up Drone

1. Set up RDS MySQL database / schema
2. Set up GitHub OAuth application for Drone
    1. See http://readme.drone.io/setup/remotes/github/
3. Update rancher-compose with:
    1. RDS database hostname
4. Update docker-compose with:
    1. `REMOTE_CONFIG`: 
        1. `client_id` and `client_secret`: from GitHub OAuth
        2. `orgs`: with private GitHub organization name
        3. `open`: [see Drone docs](http://readme.drone.io/setup/remotes/github/) (self-registration for users within org) 
    2. `DATABASE_CONFIG`: RDS credentials
    3. `PLUGIN_PARAMS`: these are global secrets that are injected into every build. It's much easier to start out this way than to use the per-repository secrets.
        1. `DOCKER_USER` and `DOCKER_PASSWORD`: CI user in Docker registry with access to publish Docker images
        2. `RANCHER_DEV_ACCESS_KEY` and `RANCHER_DEV_SECRET_KEY`:  Rancher -> API -> "Add Environment API Key". 
            1. Name: "Drone"
            2. Do this in the Dev environment. If we need Drone to deploy to other environments we can add those later.
        3. `SLACK_WEBHOOK_URL`: In case you want Slack notifications in a chat room

