// search implementation

document.addEventListener('DOMContentLoaded', function() {
    // Retrieve JSON data from script tag
    const artistsDataElement = document.getElementById('artists-data');
    const artistsData = JSON.parse(artistsDataElement.textContent);

    const input = document.getElementById('search-input');
    const suggestionsBox = document.getElementById('suggestions');

    input.addEventListener('input', () => {
        const query = input.value.toLowerCase();
        let matches = [];

        if (query) {
            matches = artistsData.filter(artist => artist.Name.toLowerCase().includes(query));
        }

        suggestionsBox.innerHTML = matches.map(match =>
            `<div class="suggestion-item">${match.Name}</div>`
        ).join('');

        // Handle clicks on suggestions
        document.querySelectorAll('.suggestion-item').forEach(item => {
            item.addEventListener('click', () => {
                input.value = item.textContent;
                suggestionsBox.innerHTML = '';
            });
        });
    });
});


// index page shuffle

// async function fetchData() {
//     const response = await fetch('https://groupietrackers.herokuapp.com/api');
//     const data = await response.json();
//     // console.log(data);
//     // return data;
// }

// function shuffleArray(array) {
//     for (let i = array.length - 1; i > 0; i--) {
//         const j = Math.floor(Math.random() * (i + 1));
//         [array[i], array[j]] = [array[j], array[i]];
//     }
// }

// function displayCards(containerId, items, type) {
//     const container = document.getElementById(containerId);
//     container.innerHTML = '';

//     items.forEach(item => {
//         const card = document.createElement('div');
//         card.className = 'card';

//         if (type === 'artist') {
//             card.innerHTML = `
//                 <img src="${item.image}" alt="${item.name}">
//                 <h4>${item.name}</h4>
//             `;
//         } else if (type === 'concert') {
//             card.innerHTML = `
//                 <h4>${item.title}</h4>
//                 <p>Date: ${item.date}</p>
//                 <p>Location: ${item.location}</p>
//             `;
//         } else if (type === 'location') {
//             card.innerHTML = `
//                 <img src="${item.image}" alt="${item.name}">
//                 <h4>${item.name}</h4>
//             `;
//         }

//         container.appendChild(card);
//     });
// }

// async function init() {
//     const data = await fetchData();

//     // Prepare artists data
//     const artists = data.artists;
//     shuffleArray(artists);
//     displayCards('artists-list', artists, 'artist');

//     // Prepare concert data
//     const concerts = data.dates.index.map(d => ({
//         title: `Concert on ${d.dates[0]}`,
//         date: d.dates[0],
//         location: data.relations.index.find(r => r.datesLocations[d.dates[0]]).datesLocations[d.dates[0]].join(', ')
//     }));
//     shuffleArray(concerts);
//     displayCards('concerts-list', concerts, 'concert');

//     // Prepare locations data
//     const locations = data.locations.index;
//     shuffleArray(locations);
//     displayCards('locations-list', locations, 'location');
// }

// init();