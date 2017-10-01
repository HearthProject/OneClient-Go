var index = {
    init: function() {
        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            index.findUser();
        })
    },
    findUser: function() {
        astilectron.send({"name": "get.user"}, function(message) {
            console.log(message)
            document.getElementById("username").innerHTML = message.payload
        })
    }
};