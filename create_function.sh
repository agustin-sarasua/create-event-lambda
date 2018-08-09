aws lambda create-function \
  --region us-east-1 \
  --function-name CreateEvent \
  --memory 128 \
  --role arn:aws:iam::161262005667:role/PeopleTrackingLambda \
  --runtime go1.x \
  --zip-file fileb:///tmp/main.zip \
  --handler main