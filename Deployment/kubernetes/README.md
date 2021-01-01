1) Make sure bank DB, and payments DB exists. Look into utilizing secure client for that operation. Details in scoretrak deployment
2) Exec inside of the container, to migrate tables and create admin user


3) kubectl create configmap payments-config --from-file=./payments-config.yml
       