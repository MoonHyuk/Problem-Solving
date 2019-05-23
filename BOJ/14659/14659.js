const fs = require('fs');
const input = fs.readFileSync('/dev/stdin').toString().split('\n');

let mountains = input[1].split(' ');
mountains.push(1000000);

let ans = mountains.reduce((most_kill, new_height) => {
    if (most_kill['height'] < parseInt(new_height)) {
        most_kill['max_kills_count'] = Math.max(most_kill['max_kills_count'], most_kill['kills_count']);
        most_kill['height'] = parseInt(new_height);
        most_kill['kills_count'] = 0;
        return most_kill;
    } else {
        most_kill['kills_count'] ++;
        return most_kill;
    }
}, {
    'height': 0,
    'kills_count': 0,
    'max_kills_count': 0
});

console.log(ans['max_kills_count']);
