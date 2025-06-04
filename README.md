# Dummy HTTP Server

➡️ Used for temporary Terraform deployments or testing

Example to use:
```
docker run -p 8080:80 tedezed/dummy-http-server:80
docker run -p 8080:8888 tedezed/dummy-http-server:8888
docker run -p 8080:3000 tedezed/dummy-http-server:3000
```

Check dummy server:
```
curl http://127.0.0.1:8080
```

Port list:
- 80
- 443
- 8000
- 8443
- 8080
- 8888
- 3000
- 5000
- 6000

### Cloud Run Example with Terraform and GitHub Actions:

Terraform code to define the API:
```
resource "google_cloud_run_v2_service" "my_awesome_api" {
  name     = "my-awesome-api"
  location = var.region

  template {
    containers {
      image = "tedezed/dummy-http-server:3000"
      ports {
        container_port = 3000
      }
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }

  lifecycle {
    # ⚠️ Ignore the image for the following CI deployments
    ignore_changes = [
      template[0].containers[0].image
    ]
  }
}
```

Now you can update the API image with your CI:
```
      - name: Authenticate to Google Cloud
        uses: google-github-actions/auth@v2
        with:
          workload_identity_provider: ${{ secrets.GCP_WORKLOAD_IDENTITY_PROVIDER }}
          service_account: ${{ secrets.GCP_SERVICE_ACCOUNT_EMAIL }}

      - name: Deploy to Cloud Run
        uses: google-github-actions/deploy-cloudrun@v2
        with:
          service: dummy-http-server
          region: us-central1
          image: ⚠️ ---> YOUR-FINAL-IMAGE <--- ⚠️
          allow-unauthenticated: true
```

Once deployed, you could modify the Terraform state to take the `latest` image from your CI instead of the dummy image.
