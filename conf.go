// Copyright 2012 <MortalSkulD@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

var defaultConfigFile = `
# Copyright 2012 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

#
# Gen KeePass CSV 1.x Format File
#
# keepass_site_name := section_name
# keepass_site_id := section["LoginName"]
# keepass_site_salt := section_name + section["SiteSalt"]
# keepass_site_url := section["WebSite"]
# keepass_site_comments := section["Comments"]
# keepass_site_pwd := pwdgen(keepass_site_id, keepass_site_salt, *encrypt_key, *encrypt_salt)
#
# output:
# *.ini -> *.keepass1x.csv
#
# http://keepass.info
#

# -----------------------------------------------------------------------------

[gmail]
LoginName = abc
WebSite = http://mail.google.com
Comments = abc@gmail.com
SiteSalt = 0x5f3759df

[taobao]
LoginName = 123
WebSite = http://www.taobao.com
Comments = 123@gmail.com
SiteSalt = 

[taobao.A]
LoginName = uer_a
WebSite = http://www.taobao.com
Comments = UserA@taobao.com
SiteSalt = 

[taobao.B]
LoginName = uer_b
WebSite = http://www.taobao.com
Comments = UserB@taobao.com
SiteSalt = 

[taobao.C]
LoginName = uer_c
WebSite = http://www.taobao.com
Comments = UserC@taobao.com
SiteSalt = 

# -----------------------------------------------------------------------------
# --encrypt_key=111 --encrypt-salt=fuckcsdn

[site]
LoginName = id3
WebSite = 
Comments = site0:N9V9FyMJ8tkScBN5
SiteSalt = 0 

[site1]
LoginName = id1
WebSite = 
Comments = 641gNmCY9YFNAQ1p
SiteSalt = 

[site2]
LoginName = id0
WebSite = 
Comments = 4HLACkWRCyDHtqtx
SiteSalt = 

[site3]
LoginName = id0
WebSite = 
Comments = 3eHtu74rMFdeRaVk
SiteSalt = 

[site4]
LoginName = id0
WebSite = 
Comments = 5DSxs623Rciz7bab
SiteSalt = 

[site5]
LoginName = id0
WebSite = 
Comments = 3cfiPrcjdrhwAgM1
SiteSalt = 

[site6]
LoginName = id0
WebSite = 
Comments = 5Las25BPXCjvtywo
SiteSalt = 

[site7]
LoginName = id0
WebSite = 
Comments = 1GK1x3GnRxLSH6DT
SiteSalt = 

[site8]
LoginName = id0
WebSite = 
Comments = 3VqQSgsRRQTeR6vL
SiteSalt = 

[site9]
LoginName = id0
WebSite = 
Comments = 5hHdKchVRPeJkFjU
SiteSalt = 

# -----------------------------------------------------------------------------

`[1:]
