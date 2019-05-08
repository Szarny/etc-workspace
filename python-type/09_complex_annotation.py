from typing import Optional, Iterable

def create_email(from_addr: str,
                 to_addr: str,
                 title: Optional[str],
                 body: str,
                 cc: Optional[Iterable[str]] = None,
                 bcc: Optional[Iterable[str]] = None) -> str:
    
    email : str = ""
    email += "FROM: {}\n".format(from_addr)
    email += "TO: {}\n".format(to_addr)
    
    if cc is not None:
        email += "CC: {}\n".format(",".join(cc))
    
    if bcc is not None:
        email += "BCC: {}\n".format(",".join(bcc))

    email += "TITLE: {}\n\n".format(title)
    email += "{}".format(body)

    return email


print(create_email("from@a.com", "to@b.com", "hello", "how r u?", ["user1@a.com", "user2@b.com"], ["user3@c.com"]))