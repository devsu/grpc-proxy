const Condor = require('condor-framework');

const Greeter = class {
  sayHello(call) {
    return { 'greeting': `Secure: Hello ${call.request.name}`};
  }
};

const options = {
  'host': 'localhost',
  'port': 3002,
  'certChain': '../ssl/localhost.pem',
  'privateKey': '../ssl/localhost.key',
};

const server = new Condor(options);
server.addService('./hello.proto', 'mysecureapp.Greeter', new Greeter());
server.start();