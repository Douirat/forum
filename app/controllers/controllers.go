package controllers

import (
    "html/template"
    "net/http"
    "forum/app/models"
)

// HomeController handles the home page
func HomePage(w http.ResponseWriter, r *http.Request) {
    // Fetch data from the model (database)
    data := models.GetDataFromDB()

    // Parse and serve the HTML template (view)
    tmpl, _ := template.ParseFiles("views/index.html")
    tmpl.Execute(w, data)
}
