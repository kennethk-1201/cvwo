const formatTime = t => {
    try {
        const arr1 = t.split("T")
        const arr2 = arr1[1].split(":")
        const arr3 = arr1[0].split("-");

        const monthMapper = {
            "01" : "January",
            "02" : "February",
            "03" : "March",
            "04" : "April",
            "05" : "May",
            "06" : "June",
            "07" : "July",
            "08" : "August",
            "09" : "September",
            "10" : "October",
            "11" : "November",
            "12" : "December"
        }

        let year = arr3[0];
        let month = monthMapper[arr3[1]];
        let day = arr3[2];

        let c = parseInt(arr2[0])

        let time = c > 12 
            ? (c - 12).toString() + ":" + arr2[1] + " p.m."
            : c.toString() + ":" + arr2[1] + " a.m."

        return "Due by " + day + " " + month + " " + year + ", " + time;
    } catch (err) {
        console.log(err);
    }
}

export default formatTime;