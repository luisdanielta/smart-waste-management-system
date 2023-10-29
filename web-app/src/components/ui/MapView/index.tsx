import Map from "react-map-gl/maplibre";
import "./styles.css";

export default function MapView() {
  return (
    <div className="map_view">
      <Map
        initialViewState={{
            /*10.422366, -75.548598*/
          longitude: -75.548598,
          latitude: 10.422366,
          zoom: 14,
        }}
        style={{ height: 730 }}
        mapStyle="https://api.maptiler.com/maps/streets/style.json?key=hpIsP2oOBBR32z1kce5w"
      />
    </div>
  );
}
