clear
echo " -----------------------------------------------------------------------------------------------------"
echo "|                                 go run . test/audit/example00.txt                                   |"
echo " -----------------------------------------------------------------------------------------------------"
time go run . test/audit/example00.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                                 go run . test/audit/example01.txt                                   |"
echo " -----------------------------------------------------------------------------------------------------"
time go run . test/audit/example01.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                                 go run . test/audit/example02.txt                                   |"
echo " -----------------------------------------------------------------------------------------------------"
time go run . test/audit/example02.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                                 go run . test/audit/example03.txt                                   |"
echo " -----------------------------------------------------------------------------------------------------"
time go run . test/audit/example03.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                                 go run . test/audit/example04.txt                                   |"
echo " -----------------------------------------------------------------------------------------------------"
time go run . test/audit/example04.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                                 go run . test/audit/example05.txt                                   |"
echo " -----------------------------------------------------------------------------------------------------"
time go run . test/audit/example05.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                               go run . test/audit/badexample00.txt                                  |"
echo " -----------------------------------------------------------------------------------------------------"
go run . test/audit/badexample00.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                               go run . test/audit/badexample01.txt                                  |"
echo " -----------------------------------------------------------------------------------------------------"
go run . test/audit/badexample01.txt
echo ""
echo ""
echo " -----------------------------------------------------------------------------------------------------"
echo "|                               go run . test/autre/badexample02.txt                                  |"
echo " -----------------------------------------------------------------------------------------------------"
go run . test/autre/badexample02.txt