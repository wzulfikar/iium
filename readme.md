# IIUM Results Visualizer

1. use offline by downloading cli program and provide json
- use online by uploading json, providing url to json or login
- dashboard contains stats of sessions passed and latest cam
- stats of sessions passed has 3 view mode: cgpa, gpa, mixed (cgpa + gpa)
- when using cgpa/gpa mode, can add up to 5 students' results (useful for comparison)
- students comparison is not available when view mode is set to mixed
- when used via login, json is available for download. eg: cgpa.iium.online/1222665 will show dashboard while cgpa.iium.online/1222665.json will show json
- provide setting for public json view or dashboard view
- js using http://morrisjs.github.io/morris.js/
- no need for db. all data stored in json. no password is not stored

# cgpa calculator
- insert current credit hours
- insert desired cgpa
- the system will display:
  - how much gpa student must get to achieve desired cgpa
  - average mark needed for each subject