cd cmd/homework-3-oguzhantasimaz/
echo "Which function do you want to execute?\n1. List\n2. Search(id || title)   \n3. Buy(id, count)\n4. Delete(id)\n"
read choice
if [ $choice -eq 1 ]
then
  echo "Listing all products"
  go run main.go list
elif [ $choice -eq 2 ]
then
  echo "Searching for a product"
  read -p "Enter the id or title of the product: " id
  go run main.go search $id
elif [ $choice -eq 3 ]
then
  echo "Buying a product"
  read -p "Enter the id of the product: " id
  read -p "Enter the count of the product: " count
  go run main.go buy $id $count
elif [ $choice -eq 4 ]
then
  echo "Deleting a product"
  read -p "Enter the id of the product: " id
  go run main.go delete $id
else
  echo "Invalid choice"
fi