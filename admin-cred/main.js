var admin = require("firebase-admin");

var serviceAccount = require("./nova-we-firebase-adminsdk-u06b9-35cee7b2d9.json");

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount)
});

var mg=admin.projectManagement();
console.log(mg);
