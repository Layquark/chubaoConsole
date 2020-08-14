# Nginx Operator

A sample Nginx Operator project for  chubaofs.


   ## Create the Operator

   ```
   operator-sdk new nginx-operator --repo github.com/chenzhiwei/nginx-operator
   operator-sdk add api --api-version operator.example.com/v1alpha1 --kind Nginx
   operator-sdk add controller --api-version operator.example.com/v1alpha1 --kind Nginx
   
   operator-sdk generate k8s
   operator-sdk generate crds
   
   operator-sdk build quay.io/siji/chubaoConsole:v0.0.1
   ```

   ### Update the generated CSV owned CRD fields:

   deploy/olm-catalog/nginx-operator/0.0.1/nginx-operator.v0.0.1.clusterserviceversion.yaml

   * description
   * displayName

   ```
   spec:
     customresourcedefinitions:
       owned:
       - kind: Nginx
         name: nginxes.operator.example.com
         version: v1alpha1
         description: The Nginx CRD
         displayName: Nginx
   ```

   ### [Optional] Update the generated CSV required CRD fields:

   ```
   spec:
     customresourcedefinitions:
       required:
       - kind: Database
         name: databases.operator.example.com
         version: v1alpha1
         description: The Database CRD
         displayName: Database
   ```

   * Makefile

   ```
   make build
   make image
   make push-image
   make push-csv
   ```

   ## References

   1. https://github.com/operator-framework/operator-sdk/blob/master/doc/user/olm-catalog/generating-a-csv.md
   1. https://github.com/operator-framework/operator-lifecycle-manager/blob/master/doc/design/building-your-csv.md