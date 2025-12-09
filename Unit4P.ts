/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview This program keeps track of car stats and simulates drive/fill/oil change
 */

// Initial data (per assignment)
const make: string = "UsedCarBrand";
const model: string = "ModelX";
let carColor: string = "Silver";
let odometer: number = 65000;       // starting mileage
let oilChangeKM: number = 65000;    // last oil change at 65000
const gasCost: number[] = new Array(10);
let fillIndex: number = 0;

// initial fill-up per assignment
gasCost[fillIndex++] = 74.0;

// Function: carStats() returns a string of car info
function carStats(): string {
  return (
    `Car Make: ${make}\n` +
    `Car Model: ${model}\n` +
    `Car Colour: ${carColor}\n` +
    `Odometer: ${odometer} km\n` +
    `Last oil change at: ${oilChangeKM} km\n` +
    `Number of recorded fill-ups: ${fillIndex}`
  );
}

// Function: wrapCar() prompts for a new colour and returns it
function wrapCar(): string {
  const newColour: string = prompt("Enter a new colour to wrap the car with: ") || carColor;
  return newColour.trim();
}

// Function: drive() returns random km driven and updates odometer
function drive(): number {
  const driven: number = Math.floor(Math.random() * 901) + 100; // 100..1000
  odometer += driven; // update odometer internally
  return driven;
}

// Function: fillUp() prompts user for the price paid and stores it in gasCost
function fillUp(): void {
  if (fillIndex >= gasCost.length) {
    console.log("Gas cost array is full; cannot record more fill-ups.");
    return;
  }
  const input = prompt("Enter how much you paid for the gas to fill up: ") || "0";
  const cost = Number(input);
  gasCost[fillIndex++] = cost;
  console.log(`Recorded fill-up of $${cost.toFixed(2)} at slot ${fillIndex - 1}.`);
}

// Function: displayCostToFillUp() shows all recorded costs and returns average
function displayCostToFillUp(): number {
  if (fillIndex === 0) {
    console.log("No fill-up records.");
    return 0;
  }
  console.log("\nRecorded fill-up costs:");
  let sum = 0;
  for (let i = 0; i < fillIndex; i++) {
    console.log(`Fill-up ${i + 1}: $${gasCost[i].toFixed(2)}`);
    sum += gasCost[i];
  }
  const avg = sum / fillIndex;
  console.log(`Average cost per fill-up: $${avg.toFixed(2)}`);
  return avg;
}

// Function: oilChange() returns boolean, prints message, and updates oilChangeKM internally
function oilChange(): boolean {
  if (odometer - oilChangeKM >= 5000) {
    oilChangeKM = odometer; // update last oil change
    console.log("\nAn oil change was done.");
    return true;
  } else {
    console.log("\nYour car does not need an oil change.");
    return false;
  }
}

/* ----------------- Main program ----------------- */

console.log("Initial car stats:");
console.log(carStats());

// wrap the car (change colour)
const colourInput = wrapCar();
carColor = colourInput;
console.log(`\nCar colour updated to: ${carColor}`);

// drive
const drivenKm = drive();
console.log(`\nYou drove ${drivenKm} km. New odometer: ${odometer} km.`);

// ask user to fill up now
fillUp();

// display fill-up costs and average
const avgCost = displayCostToFillUp();

// check oil change
oilChange();

console.log("\nFinal car stats:");
console.log(carStats());

console.log("\nDone.");
