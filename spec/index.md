# Specification
The specification is written from the mobile application perspective (frontend). It's split into the pages (views) that we can see in the app.

## The MVP idea
Collect the best skate spots around the city and make them available to skaters through a nice and shiny interface with advanced filtering options such as `user location`, `distance to a spot`, `types of obstacles`, `lighting`, `legal to skate`, and more.

The app is open-source and free to use.

## Views
- [Skate spots map](views/skate_spots_map.md) (main view)
- [Add new skate spot](views/add_new_skate_spot.md)
- [Profile](views/profile.md)
- [Auth](views/auth.md) (main view for unauthenticated users)

> **NOTE:** Every view except [Auth](views/auth.md) is protected. The user has to log in to use the app.
