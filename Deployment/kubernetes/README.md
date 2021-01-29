1) Make sure bank DB, and payments DB exists. Look into utilizing secure client for that operation. Details in scoretrak deployment

2) Deploy the backend and frontend and configure the gateway:
```
kubectl apply -f backend.yaml
kubectl apply -f frontend.yaml
kubectl apply -f istio.yaml
```
      
2) Exec inside of the container, to migrate tables and create admin user:
To exec:
```
kubectl  exec -it $(kubectl get pods | grep backend | awk 'NR==1 {print; exit}' | awk '{print $1}' ) bash
```

After, run the following to configure database: 
```
python manage.py migrate
```

To create admin user:
```
python manage.py createsuperuser
```

3) kubectl create configmap payments-config --from-file=./payments-config.yml
4) Run following to create payments cronjob:
```
kubectl apply -f payments.yaml
```
