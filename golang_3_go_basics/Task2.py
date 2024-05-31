# from datetime import datetime

# def convert_date(date):
#     parsed_date = datetime.strptime(date, "%dth %b %Y")
    
#     formatted_date = parsed_date.strftime("%Y-%m-%d")
    
#     return formatted_date

# print(convert_date("20th Oct 2052")) 
# print(convert_date("6th Jun 1933"))  
# print(convert_date("26th May 1960"))  

def convert_date(date):
    day,month,year = date.split()

    day = day[:-2]
    day = day.zfill(2)

    months = {
        "Jan":"01", "Feb":"02", "Mar":"03", "Apr":"04",
        "May":"05", "Jun":"06", "Jul":"07", "Aug":"08", 
        "Sep":"09", "Oct":"10", "Nov":"11", "Dec":"12"}
    
    month = months[month]

    format_date = f'{year}-{month}-{day}'

    return format_date

print(convert_date("20th Oct 2052")) 
print(convert_date("6th Jun 1933"))  
print(convert_date("26th May 1960"))  