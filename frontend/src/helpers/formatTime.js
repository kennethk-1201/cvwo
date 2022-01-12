const formatTime = time => {
    const arr1 = time.split("T")
    const arr2 = arr1[1].split(".")
    
    return arr1[0] + " " + arr2[0]
}

export default formatTime;