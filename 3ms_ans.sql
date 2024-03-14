SELECT Department.name AS Department , Employee.name AS Employee, Employee.salary AS Salary
FROM Employee JOIN Department ON Employee.departmentId = department.id 
WHERE
(Department.id, salary) IN (
        SELECT departmentId, MAX(salary)
        FROM Employee
        GROUP BY 1
    ); 