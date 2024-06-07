package main

func main() {
	helper, err := NewDBHelper()
	CheckError(err)
	defer helper.Close()

	// CRUD operations demonstration
	record := Record{Username: "Kanhaiya", PhoneNumber: "96610335556", City: "Munich", Address: "Germany"}

	record.City = "Mumbai"
	err = helper.updateRecord(record, "customers")
	CheckError(err)

	err = helper.deleteRecord("Shashi", "customers")
	CheckError(err)

	err = helper.insertRecord(record, "customers")
	CheckError(err)

	err = helper.getAllRecords("customers")
	CheckError(err)

	//go run main.go dbhelper.go errorhelper.go
}
