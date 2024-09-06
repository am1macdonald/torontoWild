import GeoJSON from 'ol/format/GeoJSON.js';
import Map from 'ol/Map.js';
import VectorLayer from 'ol/layer/Vector.js';
import VectorSource from 'ol/source/Vector.js';
import View from 'ol/View.js';
import {FeatureLike} from 'ol/Feature';
import {Pixel} from "ol/pixel";
import {Feature} from "ol";

const vectorLayer = new VectorLayer({
    background: '#1a2b39',
    source: new VectorSource({
        url: 'https://openlayers.org/data/vector/ecoregions.json',
        format: new GeoJSON(),
    }),
    style: {
        'fill-color': ['string', ['get', 'COLOR'], '#eee'],
    },
});

const map = new Map({
    layers: [vectorLayer],
    target: 'map',
    view: new View({
        center: [0, 0],
        zoom: 1,
    }),
});

const featureOverlay = new VectorLayer({
    source: new VectorSource(),
    map: map,
    style: {
        'stroke-color': 'rgba(255, 255, 255, 0.7)',
        'stroke-width': 2,
    },
});

let highlight: Feature | undefined;

const displayFeatureInfo = function (layer: VectorLayer<VectorSource>, pixel: Pixel) {
    if (layer === null) {
        throw new Error('layer not found');
    }
    const source = layer.getSource();
    if (source === null) {
        throw new Error('source not found');
    }
    let feature: Feature | undefined = undefined;

    map.forEachFeatureAtPixel(pixel, function (f) {
        if (f instanceof Feature) {
            feature = f;
        }
    })

    if (!feature) {
        throw new Error('feature not found');
    }

    const info = document.getElementById('info');
    if (!info) {
        throw new Error('info element not found');
    }
    if (feature) {
        info.innerHTML = feature.get('ECO_NAME') || '&nbsp;';
    } else {
        info.innerHTML = '&nbsp;';
    }

    if (feature !== highlight) {
        if (highlight) {
            source.removeFeature(highlight);
        }
        if (feature) {
            source.addFeature(feature);
        }
        highlight = feature;
    }
};

map.on('pointermove', function (evt) {
    if (evt.dragging) {
        return;
    }
    const pixel = map.getEventPixel(evt.originalEvent);
    displayFeatureInfo(featureOverlay, pixel);
});

map.on('click', function (evt) {
    displayFeatureInfo(featureOverlay, evt.pixel);
});
