import './style.css';
import Map from 'ol/Map.js';
import View from 'ol/View.js';
import TileLayer from "ol/layer/Tile";
import {OSM} from "ol/source";

const openStreetMapLayer = new TileLayer({
    source: new OSM(),
});

new Map({
    layers: [openStreetMapLayer],
    target: 'map',
    view: new View({
        center: [0, 0],
        zoom: 1,
    }),
});
