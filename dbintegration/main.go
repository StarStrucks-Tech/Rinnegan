package main

func main() {
	helper, err := NewDBHelper()
	CheckError(err)
	defer helper.Close()

	// CRUD operations demonstration
	record1 := Record{Username: "Mohit", PhoneNumber: "96610335556", City: "Munich", Address: "Antartica"}
	record2 := Record{Username: "Sohit", PhoneNumber: "96610335556", City: "Munich", Address: "Antartica"}

	record1.City = "Mumbai"
	err = helper.updateRecord(record2, "customers")
	CheckError(err)

	err = helper.deleteRecord("Rohit", "customers")
	CheckError(err)

	err = helper.insertRecord(record2, "customers")
	CheckError(err)

	err = helper.getAllRecords("customers")
	CheckError(err)

	//go run main.go constants.go dbhelper.go errorhelper.go
	//go run .
}
