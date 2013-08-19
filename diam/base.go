// Copyright 2013 Alexandre Fiori
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package diam

// BaseDict is a static Dict with a pre-loaded Base Protocol.
var BaseDict, _ = NewDict()

// BaseProtocolXML is an embedded version of the Diameter Base Protocol.
// It is always loaded when a new Dict is created by the NewDict function and
// its AVPs can be overloaded.
var BaseProtocolXML = []byte(`<?xml version="1.0" encoding="UTF-8"?>
<diameter>
  <vendor id="10415" name="3GPP"/>

  <application id="0"> <!-- Diameter Common Messages -->

    <avp name="Acct-Interim-Interval" code="85" must="M" may="P" must-not="V" encr="Y">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Accounting-Realtime-Required" code="483" must="M" may="P" must-not="V" encr="Y">
      <data type="Enumerated"/>
    </avp>

    <avp name="Acct-Multi-Session-Id" code="50" must="M" may="P" must-not="V" encr="Y">
      <data type="UTF8String"/>
    </avp>

    <avp name="Accounting-Record-Number" code="485" must="M" may="P" must-not="V" encr="Y">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Accounting-Record-Type" code="480" must="M" may="P" must-not="V" encr="Y">
      <data type="Enumerated">
        <item code="1" name="DELIVER_AND_GRANT"/>
        <item code="2" name="GRANT_AND_STORE"/>
        <item code="3" name="GRANT_AND_LOSE"/>
      </data>
    </avp>

    <avp name="Accounting-Session-Id" code="44" must="M" may="P" must-not="V" encr="Y">
      <data type="OctetString"/>
    </avp>

    <avp name="Accounting-Sub-Session-Id" code="287" must="M" may="P" must-not="V" encr="Y">
      <data type="Unsigned64"/>
    </avp>

    <avp name="Acct-Application-Id" code="259" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Auth-Application-Id" code="258" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Auth-Request-Type" code="274" must="M" may="P" must-not="V" encr="N">
      <data type="Enumerated">
        <item code="1" name="AUTHENTICATE_ONLY"/>
        <item code="2" name="AUTHORIZE_ONLY"/>
        <item code="3" name="AUTHORIZE_AUTHENTICATE"/>
      </data>
    </avp>

    <avp name="Authorization-Lifetime" code="291" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Auth-Grace-Period" code="276" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Auth-Session-State" code="277" must="M" may="P" must-not="V" encr="N">
      <data type="Enumerated">
        <item code="0" name="STATE_MAINTAINED"/>
        <item code="1" name="NO_STATE_MAINTAINED"/>
      </data>
    </avp>

    <avp name="Re-Auth-Request-Type" code="285" must="M" may="P" must-not="V" encr="N">
      <data type="Enumerated">
        <item code="0" name="AUTHORIZE_ONLY"/>
        <item code="1" name="AUTHORIZE_AUTHENTICATE"/>
      </data>
    </avp>

    <avp name="Class" code="25" must="M" may="P" must-not="V" encr="Y">
      <data type="OctetString"/>
    </avp>

    <avp name="Destination-Host" code="293" must="M" may="P" must-not="V" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Destination-Realm" code="283" must="M" may="P" must-not="V" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Disconnect-Cause" code="273" must="M" may="P" must-not="V" encr="N">
      <data type="Enumerated">
        <item code="0" name="REBOOTING"/>
        <item code="1" name="BUSY"/>
        <item code="2" name="DO_NOT_WANT_TO_TALK_TO_YOU"/>
      </data>
    </avp>

    <avp name="Error-Message" code="281" must="" may="P" must-not="V,M" encr="N">
      <data type="UTF8String"/>
    </avp>

    <avp name="Error-Reporting-Host" code="294" must="" may="P" must-not="V,M" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Event-Timestamp" code="55" must="M" may="P" must-not="V" encr="N">
      <data type="Time"/>
    </avp>

    <avp name="Experimental-Result" code="297" must="M" may="P" must-not="V" encr="N">
      <data type="Grouped">
        <avp name="Vendor-Id" code="266" must="M">
          <data type="Unsigned32"/>
        </avp>
        <avp name="Experimental-Result-Code" code="298" must="M">
          <data type="Unsigned32"/>
        </avp>
      </data>
    </avp>

    <avp name="Experimental-Result-Code" code="298" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Failed-AVP" code="279" must="M" may="P" must-not="V" encr="N">
      <data type="Grouped"/>
    </avp>

    <avp name="Firmware-Revision" code="267" must="" may="" must-not="P,V,M" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Host-IP-Address" code="257" must="M" may="P" must-not="V" encr="N">
      <data type="Address"/>
    </avp>

    <avp name="Inband-Security-Id" code="299" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Multi-Round-Time-Out" code="272" must="M" may="P" must-not="V" encr="Y">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Origin-Host" code="264" must="M" may="P" must-not="V" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Origin-Realm" code="296" must="M" may="P" must-not="V" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Origin-State-Id" code="278" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Product-Name" code="269" must="" may="" must-not="P,V,M" encr="N">
      <data type="UTF8String"/>
    </avp>

    <avp name="Proxy-Host" code="280" must="M" may="" must-not="P,V" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Proxy-Info" code="284" must="M" may="" must-not="P,V" encr="N">
      <data type="Grouped">
        <avp name="Proxy-Host" code="280" must="M">
          <data type="DiameterIdentity"/>
        </avp>
        <avp name="Proxy-State" code="33" must="M">
          <data type="OctetString"/>
        </avp>
      </data>
    </avp>

    <avp name="Proxy-State" code="33" must="M" may="" must-not="P,V" encr="N">
      <data type="OctetString"/>
    </avp>

    <avp name="Redirect-Host" code="292" must="M" may="P" must-not="V" encr="N">
      <data type="DiameterURI"/>
    </avp>

    <avp name="Redirect-Host-Usage" code="261" must="M" may="P" must-not="V" encr="N">
      <data type="Enumerated">
        <item code="0" name="DONT_CACHE"/>
        <item code="1" name="ALL_SESSION"/>
        <item code="2" name="ALL_REALM"/>
        <item code="3" name="REALM_AND_APPLICATION"/>
        <item code="4" name="ALL_APPLICATION"/>
        <item code="5" name="ALL_HOST"/>
        <item code="6" name="ALL_USER"/>
      </data>
    </avp>

    <avp name="Redirect-Max-Cache-Time" code="262" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Result-Code" code="268" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Route-Record" code="282" must="M" may="" must-not="P,V" encr="N">
      <data type="DiameterIdentity"/>
    </avp>

    <avp name="Session-Id" code="263" must="M" may="P" must-not="V" encr="Y">
      <data type="UTF8String"/>
    </avp>

    <avp name="Session-Timeout" code="27" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Session-Binding" code="270" must="M" may="P" must-not="V" encr="Y">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Session-Server-Failover" code="271" must="M" may="P" must-not="V" encr="Y">
      <data type="Enumerated">
        <item code="0" name="REFUSE_SERVICE"/>
        <item code="1" name="TRY_AGAIN"/>
        <item code="2" name="ALLOW_SERVICE"/>
        <item code="3" name="TRY_AGAIN_ALLOW_SERVICE"/>
      </data>
    </avp>

    <avp name="Supported-Vendor-Id" code="265" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Termination-Cause" code="295" must="M" may="P" must-not="V" encr="N">
      <data type="Enumerated">
        <item code="1" name="DIAMETER_LOGOUT"/>
        <item code="2" name="DIAMETER_SERVICE_NOT_PROVIDED"/>
        <item code="3" name="DIAMETER_BAD_ANSWER"/>
        <item code="4" name="DIAMETER_ADMINISTRATIVE"/>
        <item code="5" name="DIAMETER_LINK_BROKEN"/>
        <item code="6" name="DIAMETER_AUTH_EXPIRED"/>
        <item code="7" name="DIAMETER_USER_MOVED"/>
        <item code="8" name="DIAMETER_SESSION_TIMEOUT"/>
      </data>
    </avp>

    <avp name="User-Name" code="1" must="M" may="P" must-not="V" encr="Y">
      <data type="UTF8String"/>
    </avp>

    <avp name="Vendor-Id" code="266" must="M" may="P" must-not="V" encr="N">
      <data type="Unsigned32"/>
    </avp>

    <avp name="Vendor-Specific-Application-Id" code="260" must="M" may="P" must-not="V" encr="N">
      <data type="Grouped">
        <avp name="Vendor-Id" code="266" must="">
          <data type="Unsigned32"/>
        </avp>
        <avp name="Auth-Application-Id" code="258" must="M">
          <data type="Unsigned32"/>
        </avp>
        <avp name="Acct-Application-Id" code="259" must="M">
          <data type="Unsigned32"/>
        </avp>
      </data>
    </avp>

  </application>
</diameter>`)