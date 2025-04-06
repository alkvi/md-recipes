# System overview

A manually curated database of recipes in a responsive, modular, "at the fingertips" format, able to be presented in a visually slick, non-cluttered overview. Access via web browser. Self-hosted or deployed publicly.

- Recipes stored in plaintext markdown. 
- Easy import/export of recipes. 
- Recipe scraper that converts to markdown. 
- Possibly modular overview layout to user specs. By default, in a grid format with image and title as main components.

# Functional description

The user visits the main server page and logs in with username/password.
The user is greeted by an overview of all recipes. The user can 

- go to a recipe in detail and view/modify 
- create a new recipe
- create groups of recipes based on metadata tags, or drag and drop

When creating a recipe, the user is greeted with the view/modify page for a blank recipe. In this page (or from main page) the user can import a recipe from a URL or MD format. The user can also export a recipe.

The user can tag recipes with any metadata tags.

The user has user-specific settings where they can

- edit presentation of recipes in overview (rows, columns)
- edit user settings
- edit admin settings if admin (permissions etc?)

# Design considerations

## Guidelines

Very important: clean & responsive feel throughout.

Some target designs and ideas:
- Design similar to Mealie (excluding user management and other features). Main issue: clunky, proprietary, recipes are not fast to edit. Not stored in markdown or plaintext, feels like loss of ownership of data. https://hay-kot.github.io/mealie/ 
- Markdown handling similar to Chowdown. https://chowdown.io/

Markdown editor / view when browsing recipe contents.
The recipe pane is either a static view with toggle view/edit, or dynamically enters edit on press.
Likely go for dynamic editor.

## Architectural strategies

Easy to deploy self-hosted or on another host. Could scale to provide hosting in future.
Provide a nice API to retrieve recipes (content, metadata), scrape, upload recipes, etc.
Tags for categorizing recipes into groups, main ingredients, etc. Possibly in the markdown itself as a preamble.
Search function for finding certain tags, matching content.
Then just see what we want to add.

# UI Design / look

See [[md-recipes-diagram.pdf]]
Created with https://github.com/mydraft-cc/ui

# Architecture / implementation

Write something in GoLang.

Backend
https://go.dev/
https://github.com/go-chi/chi

Frontend 
https://svelte.dev/

## UI layer

Design primitives with melt-ui 
https://www.melt-ui.com/

Consume API with Svelte

Markdown parsing with markdown-it
https://github.com/markdown-it/markdown-it
## Application layer

Recipe handling logic, user handling logic

## Storage layer

Store recipes as simple plaintext markdown files.
TBD: write strategies, concurrency issues

Storage for users, settings etc in a normal SQLite DB?

# Detailed system design

In-depth description of components above if needed 

# Goals and milestones

TBD

# Timeline 

TBD

