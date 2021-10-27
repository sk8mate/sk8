# Skate spots map
**Route:** `/`

**UI Sketch**
```
###########################
. . . . . Search . . . . . 
. . . . . Filters . . . . . 
###########################
. . . . . . . . . . . . . .
. . . . . . . . . . . . . . 
. . . . . . . . . . . . . . 
. . . . . . Map . . . . . .
. . . . . . . . . . . . . . 
. . . . . . . . . . . . . . 
. . . . . Drawer . . . . . 
###########################
. . . . Navigation . . . . 
###########################
```

## User perspective
### Search & Map
As a user, I want to be able to move the map position to any place by typing the place's name in a text input to browse spots in different cities/locations.

As a user, I want the map to be centered to my current location by default.

As a user, I want to see my current location on the map.

As a user, I want to see all the spots marked on the map.

As a user, I want to be able to open a detailed view of a spot by tapping on a marker.

### Filters
As a user, I want to be able to narrow down my search to spots with specific attributes.
> Example: I'm going to skate in the late evening, so let's find spots with a `lighting`.


As a user, I want to be able to combine spots' filters.
> Example: Select `mini ramp` obstacle and `lighting` at the same time.

### Drawer
As a user, I want to be able to draw out a scrollable list of the spots with more details.

One item in a list should contain:
- photos in a horizontal view
- name
- rating
- distance
- spot's attributes

As a user, I want to get directions to the selected spot.
> Example: Click a `Navigate` button to open a map application and use the spot's address as a destination.

As a user, I want to be able to open a detailed view of a spot by tapping on a list item.

### Navigation
As a user, I want to see a navigation on every view to be able to switch between app views quickly.

**Menu items:**
- skate spots map `/`
- add new spot `/add`
- profile `/profile`

As a user, I want to be able to switch between `/add` and other views without input data loss.
> Example: I upload photos and change the view to skate spots map. I want the photos to be there when I visit the `/add` view again.