/*
  app inventory.
  takes a list of applications and places them in a database, and shows them.

  applications have many configuration items
  configuration items have a type


 */

var apps = new Vue({
    el: '#appinv',
    data: {
        appdata: 'this will soon be app data'
    },
    created: function() {
        this.loadApps();
    },
    methods: {

        loadApps: function() {

            this.message = 'loading...';
            var apps = this;

            axios.get('/applications')
                .then(function(response){
                    //apps.message = response.data[0];
                    apps.appdata =
                        response.data[1].applicationname + " " +
                        response.data[1].businessunit
                })
                .catch(function(error){
                    apps.appdata = 'An error ocurred.' + error;
                });
        }
    }
})
