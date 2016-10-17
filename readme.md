# IIUM Results Visualizer

> http://codepen.io/wzulfikar/pen/jrvVpa

1. use offline by downloading cli program and provide json
- use online by uploading json, providing url to json or login
- dashboard contains stats of sessions passed and latest cam
- stats of sessions shows cgpa, gpa & number of subjects taken
- can add up to 5 students' results for comparison purpose
- when there's more than 1 student, user must set _compare by_ option to any of these: cgpa, gpa, number of subjects taken. by default, it will cgpa will be used for the comparison.
- students comparison is not available when view mode is set to mixed
- when used via login, json is available for download. eg: cgpa.iium.online/1222665 will show dashboard while cgpa.iium.online/1222665.json will show json
- provide setting for public json view or dashboard view
- js using http://morrisjs.github.io/morris.js/
- no need for db. all data stored in json. no password is not stored
- using login is safe because we don't store your pass, and we use encrypted connection

# cgpa calculator
- insert current credit hours
- insert desired cgpa
- the system will display:
  - how much gpa student must get to achieve desired cgpa
  - average mark needed for each subject