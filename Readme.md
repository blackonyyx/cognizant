Create a simple RESTful API to manage loan of e-book in an electronic library. The API should allow users to:
1. Search for availability based on booktitle. 
2. Borrow a book.
3. Extend a book loan.
4. Return a book.


API specifications:
We do not implement a login for users. For simplicity, Users will receive a loan id that they can use to access their current loan status. Users can also use their email to find loaned books tied to them
Duration of the loan is 28 days and an extension will push the due date another 21 days. The specific timing is on epoch time. Hence adding 21 days as of the time of return.
Loans can be extended only once.
User can retrieve their Overall Receipt status by Receipt Id or Email
User can extend book due dates only BEFORE the return date of the book.

Endpoints:

"/add" : Books to be added to the library
"/book" : search for a book using partial matching of: author, title. or full matching of book id


"/borrow": All books in the borrow list must be validly available for borrowing. Otherwise the response will return a error.
"/return": Similarly all books to be extended must be valid.
Extra endpoints for flavour/ usability
/read


1. UseGoastheprogramminglanguage.
2. ImplementtheAPIwithasimplein-memorystorage(maporslice).Youwillneed
to populate some books during the start-up of the program
1. TheLibraryshouldhaveatleastthefollowingobjects.Youmayaddinanynew
objects or fields as you deemed necessary to complete the program o BookDetail
1. Title(string):Unique identifier for the book.
2. AvailableCopies(int):No of available copies of the book that can be
    loaned. 
o LoanDetail
1. Name Of Borrower(string):Nameofborrower.
2. LoanDate(date):Datewherethebookwasborrowed.
3. Return Date(date):Datewherethebookshouldbereturned.
4. Expose the following RESTful endpoints:
o GET /Book to retrieve the detail and available copies of a book title. Done
o POST /Borrow to borrow a book (loan period: 4 weeks) and display the detail of the loan. Done
o POST /Extend to extend the loan of the book (extend 3 weeks from return Done
date).
o Post /Return to return the book. Done
1. ReturnappropriateHTTPstatuscodesforsuccessanderrorscenarios.
2. WriteatleastoneunittestforeachendpointusingGo’snet/http/httptestpackage.

Relationships:
A Book is a Entity
A Receipt contains Loans that can be retrieved
A Loan is a many to 1 relation to a Book

Testing Assumptions:
Assumption of Testing is that the software will not be unreasonably adversarially used.
Testing would cover reasonable cases of use by a normal user.
Eg Adding book contents that are 100% the same.

APIs below that are for use in managing /upkeeping the entries in the library will be assumed to have correct input, as the functionality emphasis is on the business logic relating to the application. The apis for this are detailed below. Basic functionality unit testing is implemented for this but regression testing will not.

/save

The server also regularly polls itself to update the status of loans.