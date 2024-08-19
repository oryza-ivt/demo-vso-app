# demo-vso-app
Simple app to display `PodName` and `Secrets` that has been provided from `ENV`. The name of the env's should be provided as it is.

**Used `ENV` :**
- `POD_NAME` --> To display where the application deployed into.
- `SECRET_USERNAME` --> Sample secret to display
- `SECRET_PASSWORD` --> Sample secret to display

## Build Image
Sample command to build image:
(on project root directory)
```bash
docker build -t oryzaivt/demo-vso-app:latest .
```

Sample command to push Image to registry
```bash
docker push oryzaivt/demo-vso-app:latest
```


