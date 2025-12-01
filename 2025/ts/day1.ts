// NOTE: Only solved part 1
let file = (await Bun.file("./src/input/day1.txt").text()).split("\n");

let input = [
  "L68",
  "L30",
  "R48",
  "L5",
  "R60",
  "L55",
  "L1",
  "L99",
  "R14",
  "L82",
];

let current_value = 50;
let test = file.reduce((acc, n) => {
  const direction = n[0];
  let num = Number.parseInt(n.slice(1));
  if (direction == "L") {
    num = -num;
  }
  current_value = (current_value + num) % 100;
  if (current_value == 0) {
    return acc + 1;
  }
  return acc;
}, 0);

console.log(test);
