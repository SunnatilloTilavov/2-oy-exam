select employees.name AS name,departments.name AS job ,count(salaries.amount) AS count
from  employees JOIN departments ON employees.department_id=departments.id 
JOIN salaries ON employees.id=salaries.employee_id group by employees.name, departments.name;