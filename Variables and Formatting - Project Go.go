/*
Question 1
Let’s start by making a basic Go program. Before anything, we should declare our package information. In Go, this is done with the line package main. Type this as the first line of the program.

Question 2
Next we want to print out our data to the terminal, so we’ll be using the fmt package. Import fmt into our Go program.

Question 3
Now let’s create our base program by adding in a main function.

Question 4
Now let’s define the variables we’re going to use to store the data for each comic book. Comics have: a publisher, a writer, an artist, a title, a year they were released in, a grade for its condition, and a page number.

Let’s start by defining our string-type data. Define variables for publisher, writer, artist, and title. Make all of these variables strings!

Question 5 

Let’s start defining the numeric variables, we have two integers that we need to define: the year and the page number.

Define integer variables called year and pageNumber.

Question 6
We’ll need to define our condition grade. This is a score from 0.0 to 10.0 with decimal values. Because of the decimal values, we’ll need to use a float to store it.

Create a grade variable that is a float32.

Question 7 
Let’s use all of our variables in a series of fmt.Println statements! We want to output all the information from each of our comic books in a way that is repeatable, so be sure to format all of our information in a human-readable way.

We can use commas to separate string literals from variables and Println() will give us spaces between the things we input, so we can write something like:

fmt.Println(title, "written by", writer, "drawn by", artist)
Fill in the rest of the variables to give a full depiction of the comic book in question!

Question 8 

We finally have our first comic book to keep track of, let’s use software to manage it!

We’ll be adding this variable and the variables for the following steps above the print statement we previously just created. The honor of our first comic book goes to the award-winning "Mr. GoToSleep". Assign that value to title.

Click the save button to see the output change in the terminal. As we continue to assign variables their values, save the code and watch the outputs change as well!

Question 9
We need to give the author her due credit, store the author’s name "Tracey Hatchet" in our writer variable.

Question 10
Let’s keep the momentum going by storing the artist, the American legend herself "Jewel Tampson".

Question 11
Be sure to save the publisher of this graphic novel, "DizzyBooks Publishing Inc." as well.

Question 12
This copy of Mr. GoToSleep came out in 1997 so let’s use that value as our year.

Question 13
It’s a floppy comic so it’s only got 14 pages. Save that value in our pageNumber variable.

Question 14
Our beloved vintage Mr. GoToSleep comic has unfortunately seen some wear-and-tear in the intervening decades and is currently valued at a 6.5 in terms of condition. Assign that value to its grade.

Remember to save the code to see the full message with their intended values printed!

Question 15
We’re going to add another comic book, but since we want to keep the print out of first comic’s details, copy the existing print statement and paste the new statement right below existing statement.

We’ll be assigning new variables in between these two print statements to observe how values change depending on where in the code they’re located.

Question 16
Our next comic book is the esteemed first trade paperback of the incredible comic Epic. Reassign title to "Epic Vol. 1", remember add this code in between the two print statements.

Save the code to see how our output changes!

Question 17
The author, a noted comics writer in his own right, is named Ryan N. Shawn. Reassign the writer variable to give the writer his due credit.

Question 18
Reassign artist with our current artist’s name, the Spiegeler-winning "Phoebe Paperclips".

Question 19
We have the same publisher as last time "DizzyBooks Publishing Inc.", so we don’t have to reassign it here. But we can if we were trying to be diligent about having all the information in one place.

Question 20
Epic Vol. 1 was released in 2013, so use that number to update our year variable.

Question 21
As a trade paperback it’s a compilation of a few floppy comics and comes out to 160 pages. Use that value to store over our pageNumber variable.

Question 22
This comic is in fantastic condition but has a small shipping defect, warranting a 9.0 for its grade.

Question 23
Good job, you’ve gotten plenty of experience creating new variables and you’re set up to learn what’s next!

If you want to challenge yourself, consider:

Adding more variables that store a comic’s data (like the genre).
Create another comic series, with own unique set of variables.
Perform some computation using the age, grade, and pageNumber to calculate the cost of the comic book.
You can do it!
*/
package main

import "fmt"

func main() {
  var publisher, writer, artist, title, genre string
  var year, pageNumber uint
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5
  genre = "Horror"
fmt.Println(title, "written by", writer, "drawn by", artist, 
    "written in the year", year, "Published by", publisher,"with", pageNumber, "pages",
    "and the quality grade of", grade, "in the", genre,"classification")

fmt.Println("---------------------------------------------------")

// Second book entry    
  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0
  genre = "History"

fmt.Println(title, "written by", writer, "drawn by", artist, 
    "written in the year", year, "Published by", publisher,"with", pageNumber, "pages",
    "and the quality grade of", grade, "in the", genre,"classification")

fmt.Println("---------------------------------------------------")

// Custom book entry    
  title = "How to Bake Cakes from Cake Wars"
  writer = "Duff Goldman"
  artist = "Baking Company"
  publisher = "The Food Network"
  year = 2012
  pageNumber = 1060
  grade = 10.0
  genre = "Cooking"

fmt.Println(title, "written by", writer, "drawn by", artist, 
    "written in the year", year, "Published by", publisher,"with", pageNumber, "pages",
    "and the quality grade of", grade, "in the", genre,"classification")
}


















