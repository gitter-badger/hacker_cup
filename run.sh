echo "" > output.txt
date +"%T"
if [[ $1 == test ]]
then
	cat example.txt | go run main.go
else
	cat input.txt | go run main.go | tee output.txt
fi
date +"%T"
echo "DONE"
