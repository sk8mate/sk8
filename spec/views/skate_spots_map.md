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
. . . . . . . . . . . . . . 
###########################
. . . . Navigation . . . . 
###########################
```

## Search
By default, the map centers itself on the user's location. However, the map position can be quickly changed, for example, by typing another city name.

**Frontside**
- autocomplete input
- make throttled API calls while typing
- move [the map](#map) to the selected area (on submission)

**Backside**
- `/places GET`
  - `search` string
- return `X` matching places based on the `search` parameter value
```go
type Coordinates struct {
  lat float64
  long float64
}

type Place struct {
  Coordinates
  name string
  address string
}
```

## Filters
Users can filter the skate spots by their parameters, such as `lighting` or `types of obstacle`. 

**Frontside**

**Backside**
## Map

## Navigation