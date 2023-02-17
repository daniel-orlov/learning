// I made it into a terrible one-liner on purpose
function whosPaying(names) {
    return `${names[Math.floor(Math.random() * names.length)]} is going to buy lunch today`;
}

console.log(whosPaying(["Angela", "Ben", "Jenny", "Michael", "Chloe"]));