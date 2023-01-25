
function closestNumber(numbers) {
    let difference = numbers[1] - numbers[0]
    for (let i = 0; i < numbers.length; i++) {
        if (numbers[i + 1] - (numbers[i]) <= difference) {
            difference = numbers[i] - numbers[i + 1]
        }
        console.log(numbers[i])
        console.log(difference)
    }
    console.log(difference)
}
const numbers = [10, 5, 3, 1]
closestNumber(numbers)