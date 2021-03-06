/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */
package generated

import "github.com/deckarep/golang-set"

const VoipConfigDataClassId ClassID = ClassID(138)

var voipconfigdataBME *ManagedEntityDefinition

// VoipConfigData (class ID #138)
//	The VoIP configuration data ME defines the configuration for VoIP in the ONU. The OLT uses this
//	ME to discover the VoIP signalling protocols and configuration methods supported by this ONU.
//	The OLT then uses this ME to select the desired signalling protocol and configuration method.
//	The entity is conditionally required for ONUs that offer VoIP services.
//
//	An ONU that supports VoIP services automatically creates an instance of this ME.
//
//	Relationships
//		One instance of this ME is associated with the ONU.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. There is only
//			one instance, number 0. (R) (mandatory) (2 bytes)
//
//		Available Signalling Protocols
//			(R) (mandatory) (1 byte)
//
//		Signalling Protocol Used
//			(R, W) (mandatory) (1 byte)
//
//		Available Voip Configuration Methods
//			Bits 5..24 are reserved by ITU-T. Bits 25..32 are reserved for proprietary vendor configuration
//			capabilities. (R) (mandatory) (4 bytes)
//
//		Voip Configuration Method Used
//			(R, W) (mandatory) (1 byte)
//
//		Voip Configuration Address Pointer
//			The default value is 0xFFFF (R, W) (mandatory) (2 bytes)
//
//		Voip Configuration State
//			Other values are reserved. At ME instantiation, the ONU sets this attribute to 0. (R)
//			(mandatory) (1 byte)
//
//		Retrieve Profile
//			Retrieve profile: This attribute provides a means by which the ONU may be notified that a new
//			VoIP profile should be retrieved. By setting this attribute, the OLT triggers the ONU to
//			retrieve a new profile. The actual value in the set action is ignored because it is the action
//			of setting that is important. (W) (mandatory) (1 byte)
//
//		Profile Version
//			Profile version: This attribute is a character string that identifies the version of the last
//			retrieved profile. (R) (mandatory) (25 bytes)
//
type VoipConfigData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	voipconfigdataBME = &ManagedEntityDefinition{
		Name:    "VoipConfigData",
		ClassID: 138,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0XFF00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", 0, mapset.NewSetWith(Read), false, false, false, false, 0),
			1: ByteField("AvailableSignallingProtocols", 0, mapset.NewSetWith(Read), false, false, false, false, 1),
			2: ByteField("SignallingProtocolUsed", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 2),
			3: Uint32Field("AvailableVoipConfigurationMethods", 0, mapset.NewSetWith(Read), false, false, false, false, 3),
			4: ByteField("VoipConfigurationMethodUsed", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 4),
			5: Uint16Field("VoipConfigurationAddressPointer", 0, mapset.NewSetWith(Read, Write), false, false, false, false, 5),
			6: ByteField("VoipConfigurationState", 0, mapset.NewSetWith(Read), false, false, false, false, 6),
			7: ByteField("RetrieveProfile", 0, mapset.NewSetWith(Write), false, false, false, false, 7),
			8: MultiByteField("ProfileVersion", 25, nil, mapset.NewSetWith(Read), true, false, false, false, 8),
		},
	}
}

// NewVoipConfigData (class ID 138 creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from the wire, about to be sent on the wire.
func NewVoipConfigData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(voipconfigdataBME, params...)
}
