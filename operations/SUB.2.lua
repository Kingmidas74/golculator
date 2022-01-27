if args[3] == nil or args[3] == '' then
    result[1]=args[1] * (-1)
    result[2]=args[2] * (-1)
    return
end
result[1]=args[1]-args[3]
result[2]=args[2]-args[4]