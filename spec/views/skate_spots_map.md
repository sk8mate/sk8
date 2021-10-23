# Skate spots map
**Route:** `/`

The main view of the app.

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

## Search
By default, the map centers itself on the user's location. However, the map position can be quickly changed, for example, by typing another city name.

**Frontside**
- autocomplete input
- make throttled API calls while typing
- move [the Map](#map) to the selected area (on submission)

**Backside**
- `/places GET`
  Return `X` matching places based on the `search` parameter value.
  
  Parameters
  | name     | type   | description                                   | required |
  | -------- | ------ | --------------------------------------------- | -------- |
  | search   | string | any point of interest (e.g. city or district) | yes      |
  | language | string | user's device language                        | yes      |

  **Proto**
  [https://github.com/sk8mate/sk8/tree/main/backside/places/proto](https://github.com/sk8mate/sk8/tree/main/backside/places/proto)

## Filters
Users can filter the skate spots by their parameters, such as `lighting` or `types of obstacle`. 

**Frontside**
- require data from `/spots/filters` endpoint to render
- nice buttons with checkbox functionality under the hood
- scrollable, stacked in one row, can go beyond the screen
- trigger `/spots` API request on every filter change

**Backside**
- `/spots/filters GET`
  Return list of available filters.

- `/spots GET`
  Return list of spots based on the input location (lat, long params).
  
  Parameters:
  <table>
    <tr>
      <td><strong>name</strong></td>
      <td><strong>type<strong></td>
      <td><strong>description</strong></td>
      <td><strong>required</strong></td>
    </tr>
    <tr>
      <td>lat</td>
      <td>float</td>
      <td rowspan="2">center point on the map</td>
      <td>yes</td>
    </tr>
    <tr>
      <td>long</td>
      <td>float</td>
      <td>yes</td>
    </tr>
    <tr>
      <td colspan="4">### Filters ###</td>
    </tr>
    <tr>
      <td>lighting</td>
      <td>bool</td>
      <td>if a place is lit</td>
      <td>no</td>

    </tr>
    <tr>
      <td>friendly</td>
      <td>bool</td>
      <td>if a place is intended for skaters (e.g. skatepark), not a place in front of a some kind of building where guards are trying to get rid of you</td>
      <td>no</td>
    </tr>
  </table>
## Map
**Frontside**
- center at the current user location by default
- skate spots markers
- spot details after click (to be described)
  
**Backside**
Use `/spots GET` described in [Filters](#filters).

## Drawer
Drawer with a list view of filtered spots including some more details.
**Frontside**
It should use the data from `/spots GET`.

One item in a list should contain:
- photos in a horizontal view
- name
- rating
- distance
- spot details

## Navigation
**Frontside**
- visible on every view
- even spacing between menu items
- highlight the current view
