.globl _main
.extern _printf
.extern _scanf

.section __TEXT,__cstring
msg1:      .asciz "Введите количество символов N: "
msg2:      .asciz "Количество чисел у которых сумма больше 10: %d\n"
err:       .asciz "Ошибка ввода\n"
fmt_in:    .asciz "%d"

.section __TEXT,__text

_main:
    stp x29, x30, [sp, -64]!
    mov x29, sp

    // printf("Введите количество символов N: ");
    adrp x0, msg1@PAGE
    add  x0, x0, msg1@PAGEOFF
    bl _printf

    // scanf("%d", &N);
    adrp x0, fmt_in@PAGE
    add  x0, x0, fmt_in@PAGEOFF
    add  x9, sp, #28          // &N
    str  x9, [sp, #0]         // variadic arg -> stack
    bl _scanf

    cmp w0, #1
    bne error

    ldr w8, [sp, #28]         // N
    cmp w8, #0
    ble error

    mov w9, #0                // count = 0
    str w9, [sp, #24]

    mov w10, #0               // i = 0
    str w10, [sp, #20]

loop:
    ldr w10, [sp, #20]        // i
    ldr w8,  [sp, #28]        // N
    cmp w10, w8
    bge done

    // scanf("%d", &x);
    adrp x0, fmt_in@PAGE
    add  x0, x0, fmt_in@PAGEOFF
    add  x9, sp, #16          // &x
    str  x9, [sp, #0]         // variadic arg -> stack
    bl _scanf

    cmp w0, #1
    bne error

    ldr w11, [sp, #16]        // x
    cmp w11, #0
    ble error

    mov w12, #0               // sum = 0

sum_loop:
    cmp w11, #0
    beq sum_done

    mov  w13, #10
    udiv w14, w11, w13
    msub w15, w14, w13, w11   // w15 = w11 - w14*10
    add  w12, w12, w15
    mov  w11, w14
    b sum_loop

sum_done:
    cmp w12, #10
    ble next

    ldr w9, [sp, #24]
    add w9, w9, #1
    str w9, [sp, #24]

next:
    ldr w10, [sp, #20]
    add w10, w10, #1
    str w10, [sp, #20]
    b loop

done:
    // printf("Количество чисел у которых сумма цифр больше 10: %d\n", count);
    adrp x0, msg2@PAGE
    add  x0, x0, msg2@PAGEOFF
    ldr  w9, [sp, #24]
    str  x9, [sp, #0]         // variadic arg -> stack
    bl _printf

    mov w0, #0
    ldp x29, x30, [sp], #64
    ret

error:
    adrp x0, err@PAGE
    add  x0, x0, err@PAGEOFF
    bl _printf

    mov w0, #0
    ldp x29, x30, [sp], #64
    ret