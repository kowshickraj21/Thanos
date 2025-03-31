section .text 
global _start 
_start: 
mov rax, 0X0A3831 ;
push rax ;
mov rdx, 3 ;
call print 
pop rax ;
mov rax, 60 ;
mov rdi, 20 ;
syscall ;
print: 
mov rax, 1 ;
mov rdi, 1 ;
lea rsi, [rsp+8] ;
syscall ;
ret ;
