# How to

![infra instance with multiple network](./img/03-multiple-network.png "infra instance with multiple network")

### Create stack

```
terraform apply
```

This script will create:
-   1 vpc
-   2 networks
-   2 instances http
-   3 instances db

### Delete stack

```
terraform destroy
```

### Create environment variables

1. Using CircleCI web app, click on **Organization Settings > Contexts** on the left side navigation.
2. Create two context with names "context_development" and "context_main".
3. Choose "context_development".Click on **Add Enviroment Variable** for test environment:

| Environment Variable Name   |Value|
| --- | --- |
| AWS_ACCESS_KEY_ID          | <TEST_AWS_ACCESS_KEY_ID> |
| AWS_SECRET_ACCESS_KEY      | <TEST_AWS_SECRET_ACCESS_KEY> |
| BUCKET_NAME                | <TEST_BUCKET_NAME> |
| REGION                     | <TEST_BUCKET_REGION> |

4. Choose "context_main" and add the variables for production environment.

| Environment Variable Name   |Value|
| --- | --- |
| AWS_ACCESS_KEY_ID          | <PRODUCTION_AWS_ACCESS_KEY_ID> |
| AWS_SECRET_ACCESS_KEY      | <PRODUCTION_AWS_SECRET_ACCESS_KEY> |
| BUCKET_NAME                | <PRODUCTION_BUCKET_NAME> |
| REGION                     | <PRODUCTION_BUCKET_REGION> |

5. More about [Contexts](https://circleci.com/docs/env-vars/#contexts)
     
