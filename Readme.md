Create a simple RESTful API to manage loan of e-book in an electronic library. The API should allow users to:
1. Search for availability based on booktitle. 
2. Borrow a book.
3. Extend a book loan.
4. Return a book.


API specifications:
For simplicity, Users will receive a loan id that they can use to access their current loan status. Users can also use their email to find loaned books tied to them
Duration of the loan is 28 days and an extension will push the due date another 21 days. The specific timing is on epoch time. Hence adding 21 days as of the time of return.
Loans can be extended only once.
User can retrieve their Overall Receipt status by Receipt Id or Email
User can extend book due dates only BEFORE the return date of the book.

Endpoints:

"/search" : query params id, (partial matching): title, author
Search fields are a conjunction of partial matching.

"/book" : search for a book using partial matching of: author, title. or full matching of book id
"/borrow": All books in the borrow list must be validly available for borrowing (with valid stock). Otherwise the response will return a error.
"/return": Similarly all books to be extended must be valid.
"/extend": Extends books given email + loan ids

Extra endpoints for flavour/ usability 
"/add" : Books to be added to the library (for convenience of population. no testing)
"/read": Returns the contents of the book

Relationships:
A Book is a Entity
A Receipt contains Loans that can be retrieved
A Receipt contains Loans of books
1 Book Loan = 1 stock of the book

Testing Assumptions:
Assumption of Testing is that the software will not be unreasonably adversarially used.
Testing would cover reasonable cases of use by a normal user.
Eg Adding book contents that are 100% the same.

API below that are for use in managing /upkeeping the entries in the library will be assumed to have correct input, as the functionality emphasis is on the business logic relating to the user facing application.

Optional Tasks Completed:
Implement logging for API requests and responses.
Add Validation for missing/invalid input

Addendeum:
1. Standardise Error Format, Prompt, Internal Logs, Context Messages
2. Swagger documentation added in openapi.yaml
3. Remove unused.
4. Added postman collection.


Quickstart:
```
go run main.go
```
Unit Testing:
```
go run 
```


Access the endpoints via POSTMAN collection.

```
./Library.postman_collection.json
```

Run:
1. Add 1
2. Add 2
3. Add 3

Acceptance testing of other api endpoints can begin.
