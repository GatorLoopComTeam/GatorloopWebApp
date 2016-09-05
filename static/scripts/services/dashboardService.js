angular.module('gatorloopWebApp')

.service('dashboardService', function($http) {
    var api = "/api/";
    return {
        sendStopSignal: function() {
            return $http.get(api + "sendStopSignal");
        },
        getData: function(url) {
            return $http.get(api + url);
        },
        postData: function(url, data) {
            console.log("Posting data: " + data + " to url " + url);
            return $http.post(api + url, data);
        }
    }
})
