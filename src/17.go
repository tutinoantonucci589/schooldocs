package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/GoogleCloudPlatform/go-schooldocs/school/docs/model"
)

func main() {
    // Set random seed
    time.Sleep(rand.Intn(1000))
    rand.Seed(time.Now().UnixNano())

    // Read input from user or command line
    fmt.Println("Enter a document ID:")
    docId := readInt()
    if docId < 0 || docId >= 10 {
        fmt.Println("Invalid document ID. Please enter a valid number.")
        return
    }

    // Get the content of the document using the provided input
    document, err := model.ReadFile(docId)
    if err != nil {
        fmt.Printf("Error reading document %d: %s\n", docId, err.Error())
        return
    }
    if len(document) == 0 || len(document.Title) <= 0 {
        fmt.Println("Document not found or empty.")
        return
    }

    // Print the content of the document
    fmt.Printf("Document ID: %d\n", docId)
    fmt.Printf("Title: %s\n", document.Title)
    fmt.Printf("Author: %s\n", document.Author)
    fmt.Printf("Content:\n")
    fmt.Println(document.Content)

    // Example function to check if a given author exists in the document
    checkAuthorExists := func(authorName string) bool {
        return isAuthorInDocument(document, authorName)
    }

    // Get the name of the current author
    currentAuthorName := getAuthorName()

    // Print the name of the current author
    fmt.Printf("Current Author Name: %s\n", currentAuthorName)

    // Output a message indicating which author or authors are present in the document
    if checkAuthorExists(currentAuthorName) {
        fmt.Println("The following authors are known to be present:")
        for _, author := range getAuthors() {
            fmt.Printf("%s\n", author)
        }
    } else {
        fmt.Println("No known authors.")
    }

    // Function to add an author to the document
    addAuthor := func(authorName, authorId string) {
        docIdToSave := 0

        for _, authorID := range getAuthors() {
            if isAuthorInDocument(document, authorName) {
                continue
            }
            docIdToSave = authorID
            break
        }

        // Add the new author to the document
        newDoc := model.Document{
            Title:     string(authorName),
            Author:    authorId,
            Content:   "",
            Id:        docId + 10, // Incrementing the ID for better uniqueness
            DocumentID: "doc-1",
        }

        if err = model.WriteDocument(newDoc); err != nil {
            fmt.Printf("Error writing document %d: %s\n", newDoc.DocumentID, err.Error())
            return
        }

        currentAuthorName = getAuthorName()
        docId += 10

        for _, author := range getAuthors() {
            if isAuthorInDocument(docId + 10 - 1, author.Name) {
                continue
            }
            docIdToSave = docId
            break
        }

        // Add the new author to the document and update its ID
        newDoc.Author = authorName
        docId += 10

        if err := model.WriteDocument(newDoc); err != nil {
            fmt.Printf("Error writing document %d: %s\n", newDoc.DocumentID, err.Error())
            return
        }

        currentAuthorName = getAuthorName()
    }
