rm index.zip 
cd lambda
zip -X -r index.zip *
cp index.zip ../index.zip
rm index.zip
cd ..
aws lambda update-function-code --function-name CreateEvent --zip-file fileb://index.zip