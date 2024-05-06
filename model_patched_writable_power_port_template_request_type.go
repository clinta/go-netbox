/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.7.7 (3.7)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// PatchedWritablePowerPortTemplateRequestType * `iec-60320-c6` - C6 * `iec-60320-c8` - C8 * `iec-60320-c14` - C14 * `iec-60320-c16` - C16 * `iec-60320-c20` - C20 * `iec-60320-c22` - C22 * `iec-60309-p-n-e-4h` - P+N+E 4H * `iec-60309-p-n-e-6h` - P+N+E 6H * `iec-60309-p-n-e-9h` - P+N+E 9H * `iec-60309-2p-e-4h` - 2P+E 4H * `iec-60309-2p-e-6h` - 2P+E 6H * `iec-60309-2p-e-9h` - 2P+E 9H * `iec-60309-3p-e-4h` - 3P+E 4H * `iec-60309-3p-e-6h` - 3P+E 6H * `iec-60309-3p-e-9h` - 3P+E 9H * `iec-60309-3p-n-e-4h` - 3P+N+E 4H * `iec-60309-3p-n-e-6h` - 3P+N+E 6H * `iec-60309-3p-n-e-9h` - 3P+N+E 9H * `iec-60906-1` - IEC 60906-1 * `nbr-14136-10a` - 2P+T 10A (NBR 14136) * `nbr-14136-20a` - 2P+T 20A (NBR 14136) * `nema-1-15p` - NEMA 1-15P * `nema-5-15p` - NEMA 5-15P * `nema-5-20p` - NEMA 5-20P * `nema-5-30p` - NEMA 5-30P * `nema-5-50p` - NEMA 5-50P * `nema-6-15p` - NEMA 6-15P * `nema-6-20p` - NEMA 6-20P * `nema-6-30p` - NEMA 6-30P * `nema-6-50p` - NEMA 6-50P * `nema-10-30p` - NEMA 10-30P * `nema-10-50p` - NEMA 10-50P * `nema-14-20p` - NEMA 14-20P * `nema-14-30p` - NEMA 14-30P * `nema-14-50p` - NEMA 14-50P * `nema-14-60p` - NEMA 14-60P * `nema-15-15p` - NEMA 15-15P * `nema-15-20p` - NEMA 15-20P * `nema-15-30p` - NEMA 15-30P * `nema-15-50p` - NEMA 15-50P * `nema-15-60p` - NEMA 15-60P * `nema-l1-15p` - NEMA L1-15P * `nema-l5-15p` - NEMA L5-15P * `nema-l5-20p` - NEMA L5-20P * `nema-l5-30p` - NEMA L5-30P * `nema-l5-50p` - NEMA L5-50P * `nema-l6-15p` - NEMA L6-15P * `nema-l6-20p` - NEMA L6-20P * `nema-l6-30p` - NEMA L6-30P * `nema-l6-50p` - NEMA L6-50P * `nema-l10-30p` - NEMA L10-30P * `nema-l14-20p` - NEMA L14-20P * `nema-l14-30p` - NEMA L14-30P * `nema-l14-50p` - NEMA L14-50P * `nema-l14-60p` - NEMA L14-60P * `nema-l15-20p` - NEMA L15-20P * `nema-l15-30p` - NEMA L15-30P * `nema-l15-50p` - NEMA L15-50P * `nema-l15-60p` - NEMA L15-60P * `nema-l21-20p` - NEMA L21-20P * `nema-l21-30p` - NEMA L21-30P * `nema-l22-30p` - NEMA L22-30P * `cs6361c` - CS6361C * `cs6365c` - CS6365C * `cs8165c` - CS8165C * `cs8265c` - CS8265C * `cs8365c` - CS8365C * `cs8465c` - CS8465C * `ita-c` - ITA Type C (CEE 7/16) * `ita-e` - ITA Type E (CEE 7/6) * `ita-f` - ITA Type F (CEE 7/4) * `ita-ef` - ITA Type E/F (CEE 7/7) * `ita-g` - ITA Type G (BS 1363) * `ita-h` - ITA Type H * `ita-i` - ITA Type I * `ita-j` - ITA Type J * `ita-k` - ITA Type K * `ita-l` - ITA Type L (CEI 23-50) * `ita-m` - ITA Type M (BS 546) * `ita-n` - ITA Type N * `ita-o` - ITA Type O * `usb-a` - USB Type A * `usb-b` - USB Type B * `usb-c` - USB Type C * `usb-mini-a` - USB Mini A * `usb-mini-b` - USB Mini B * `usb-micro-a` - USB Micro A * `usb-micro-b` - USB Micro B * `usb-micro-ab` - USB Micro AB * `usb-3-b` - USB 3.0 Type B * `usb-3-micro-b` - USB 3.0 Micro B * `dc-terminal` - DC Terminal * `saf-d-grid` - Saf-D-Grid * `neutrik-powercon-20` - Neutrik powerCON (20A) * `neutrik-powercon-32` - Neutrik powerCON (32A) * `neutrik-powercon-true1` - Neutrik powerCON TRUE1 * `neutrik-powercon-true1-top` - Neutrik powerCON TRUE1 TOP * `ubiquiti-smartpower` - Ubiquiti SmartPower * `hardwired` - Hardwired * `other` - Other
type PatchedWritablePowerPortTemplateRequestType string

// List of PatchedWritablePowerPortTemplateRequest_type
const (
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60320_C6               PatchedWritablePowerPortTemplateRequestType = "iec-60320-c6"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60320_C8               PatchedWritablePowerPortTemplateRequestType = "iec-60320-c8"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60320_C14              PatchedWritablePowerPortTemplateRequestType = "iec-60320-c14"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60320_C16              PatchedWritablePowerPortTemplateRequestType = "iec-60320-c16"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60320_C20              PatchedWritablePowerPortTemplateRequestType = "iec-60320-c20"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60320_C22              PatchedWritablePowerPortTemplateRequestType = "iec-60320-c22"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_P_N_E_4H         PatchedWritablePowerPortTemplateRequestType = "iec-60309-p-n-e-4h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_P_N_E_6H         PatchedWritablePowerPortTemplateRequestType = "iec-60309-p-n-e-6h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_P_N_E_9H         PatchedWritablePowerPortTemplateRequestType = "iec-60309-p-n-e-9h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_2P_E_4H          PatchedWritablePowerPortTemplateRequestType = "iec-60309-2p-e-4h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_2P_E_6H          PatchedWritablePowerPortTemplateRequestType = "iec-60309-2p-e-6h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_2P_E_9H          PatchedWritablePowerPortTemplateRequestType = "iec-60309-2p-e-9h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_3P_E_4H          PatchedWritablePowerPortTemplateRequestType = "iec-60309-3p-e-4h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_3P_E_6H          PatchedWritablePowerPortTemplateRequestType = "iec-60309-3p-e-6h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_3P_E_9H          PatchedWritablePowerPortTemplateRequestType = "iec-60309-3p-e-9h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_3P_N_E_4H        PatchedWritablePowerPortTemplateRequestType = "iec-60309-3p-n-e-4h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_3P_N_E_6H        PatchedWritablePowerPortTemplateRequestType = "iec-60309-3p-n-e-6h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60309_3P_N_E_9H        PatchedWritablePowerPortTemplateRequestType = "iec-60309-3p-n-e-9h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_IEC_60906_1                PatchedWritablePowerPortTemplateRequestType = "iec-60906-1"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NBR_14136_10A              PatchedWritablePowerPortTemplateRequestType = "nbr-14136-10a"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NBR_14136_20A              PatchedWritablePowerPortTemplateRequestType = "nbr-14136-20a"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_1_15P                 PatchedWritablePowerPortTemplateRequestType = "nema-1-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_5_15P                 PatchedWritablePowerPortTemplateRequestType = "nema-5-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_5_20P                 PatchedWritablePowerPortTemplateRequestType = "nema-5-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_5_30P                 PatchedWritablePowerPortTemplateRequestType = "nema-5-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_5_50P                 PatchedWritablePowerPortTemplateRequestType = "nema-5-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_6_15P                 PatchedWritablePowerPortTemplateRequestType = "nema-6-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_6_20P                 PatchedWritablePowerPortTemplateRequestType = "nema-6-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_6_30P                 PatchedWritablePowerPortTemplateRequestType = "nema-6-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_6_50P                 PatchedWritablePowerPortTemplateRequestType = "nema-6-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_10_30P                PatchedWritablePowerPortTemplateRequestType = "nema-10-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_10_50P                PatchedWritablePowerPortTemplateRequestType = "nema-10-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_14_20P                PatchedWritablePowerPortTemplateRequestType = "nema-14-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_14_30P                PatchedWritablePowerPortTemplateRequestType = "nema-14-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_14_50P                PatchedWritablePowerPortTemplateRequestType = "nema-14-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_14_60P                PatchedWritablePowerPortTemplateRequestType = "nema-14-60p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_15_15P                PatchedWritablePowerPortTemplateRequestType = "nema-15-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_15_20P                PatchedWritablePowerPortTemplateRequestType = "nema-15-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_15_30P                PatchedWritablePowerPortTemplateRequestType = "nema-15-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_15_50P                PatchedWritablePowerPortTemplateRequestType = "nema-15-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_15_60P                PatchedWritablePowerPortTemplateRequestType = "nema-15-60p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L1_15P                PatchedWritablePowerPortTemplateRequestType = "nema-l1-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L5_15P                PatchedWritablePowerPortTemplateRequestType = "nema-l5-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L5_20P                PatchedWritablePowerPortTemplateRequestType = "nema-l5-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L5_30P                PatchedWritablePowerPortTemplateRequestType = "nema-l5-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L5_50P                PatchedWritablePowerPortTemplateRequestType = "nema-l5-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L6_15P                PatchedWritablePowerPortTemplateRequestType = "nema-l6-15p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L6_20P                PatchedWritablePowerPortTemplateRequestType = "nema-l6-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L6_30P                PatchedWritablePowerPortTemplateRequestType = "nema-l6-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L6_50P                PatchedWritablePowerPortTemplateRequestType = "nema-l6-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L10_30P               PatchedWritablePowerPortTemplateRequestType = "nema-l10-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L14_20P               PatchedWritablePowerPortTemplateRequestType = "nema-l14-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L14_30P               PatchedWritablePowerPortTemplateRequestType = "nema-l14-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L14_50P               PatchedWritablePowerPortTemplateRequestType = "nema-l14-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L14_60P               PatchedWritablePowerPortTemplateRequestType = "nema-l14-60p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L15_20P               PatchedWritablePowerPortTemplateRequestType = "nema-l15-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L15_30P               PatchedWritablePowerPortTemplateRequestType = "nema-l15-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L15_50P               PatchedWritablePowerPortTemplateRequestType = "nema-l15-50p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L15_60P               PatchedWritablePowerPortTemplateRequestType = "nema-l15-60p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L21_20P               PatchedWritablePowerPortTemplateRequestType = "nema-l21-20p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L21_30P               PatchedWritablePowerPortTemplateRequestType = "nema-l21-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEMA_L22_30P               PatchedWritablePowerPortTemplateRequestType = "nema-l22-30p"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_CS6361C                    PatchedWritablePowerPortTemplateRequestType = "cs6361c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_CS6365C                    PatchedWritablePowerPortTemplateRequestType = "cs6365c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_CS8165C                    PatchedWritablePowerPortTemplateRequestType = "cs8165c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_CS8265C                    PatchedWritablePowerPortTemplateRequestType = "cs8265c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_CS8365C                    PatchedWritablePowerPortTemplateRequestType = "cs8365c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_CS8465C                    PatchedWritablePowerPortTemplateRequestType = "cs8465c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_C                      PatchedWritablePowerPortTemplateRequestType = "ita-c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_E                      PatchedWritablePowerPortTemplateRequestType = "ita-e"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_F                      PatchedWritablePowerPortTemplateRequestType = "ita-f"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_EF                     PatchedWritablePowerPortTemplateRequestType = "ita-ef"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_G                      PatchedWritablePowerPortTemplateRequestType = "ita-g"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_H                      PatchedWritablePowerPortTemplateRequestType = "ita-h"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_I                      PatchedWritablePowerPortTemplateRequestType = "ita-i"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_J                      PatchedWritablePowerPortTemplateRequestType = "ita-j"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_K                      PatchedWritablePowerPortTemplateRequestType = "ita-k"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_L                      PatchedWritablePowerPortTemplateRequestType = "ita-l"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_M                      PatchedWritablePowerPortTemplateRequestType = "ita-m"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_N                      PatchedWritablePowerPortTemplateRequestType = "ita-n"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_ITA_O                      PatchedWritablePowerPortTemplateRequestType = "ita-o"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_A                      PatchedWritablePowerPortTemplateRequestType = "usb-a"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_B                      PatchedWritablePowerPortTemplateRequestType = "usb-b"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_C                      PatchedWritablePowerPortTemplateRequestType = "usb-c"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_MINI_A                 PatchedWritablePowerPortTemplateRequestType = "usb-mini-a"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_MINI_B                 PatchedWritablePowerPortTemplateRequestType = "usb-mini-b"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_MICRO_A                PatchedWritablePowerPortTemplateRequestType = "usb-micro-a"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_MICRO_B                PatchedWritablePowerPortTemplateRequestType = "usb-micro-b"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_MICRO_AB               PatchedWritablePowerPortTemplateRequestType = "usb-micro-ab"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_3_B                    PatchedWritablePowerPortTemplateRequestType = "usb-3-b"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_USB_3_MICRO_B              PatchedWritablePowerPortTemplateRequestType = "usb-3-micro-b"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_DC_TERMINAL                PatchedWritablePowerPortTemplateRequestType = "dc-terminal"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_SAF_D_GRID                 PatchedWritablePowerPortTemplateRequestType = "saf-d-grid"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_20        PatchedWritablePowerPortTemplateRequestType = "neutrik-powercon-20"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_32        PatchedWritablePowerPortTemplateRequestType = "neutrik-powercon-32"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_TRUE1     PatchedWritablePowerPortTemplateRequestType = "neutrik-powercon-true1"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_NEUTRIK_POWERCON_TRUE1_TOP PatchedWritablePowerPortTemplateRequestType = "neutrik-powercon-true1-top"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_UBIQUITI_SMARTPOWER        PatchedWritablePowerPortTemplateRequestType = "ubiquiti-smartpower"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_HARDWIRED                  PatchedWritablePowerPortTemplateRequestType = "hardwired"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_OTHER                      PatchedWritablePowerPortTemplateRequestType = "other"
	PATCHEDWRITABLEPOWERPORTTEMPLATEREQUESTTYPE_EMPTY                      PatchedWritablePowerPortTemplateRequestType = ""
)

// All allowed values of PatchedWritablePowerPortTemplateRequestType enum
var AllowedPatchedWritablePowerPortTemplateRequestTypeEnumValues = []PatchedWritablePowerPortTemplateRequestType{
	"iec-60320-c6",
	"iec-60320-c8",
	"iec-60320-c14",
	"iec-60320-c16",
	"iec-60320-c20",
	"iec-60320-c22",
	"iec-60309-p-n-e-4h",
	"iec-60309-p-n-e-6h",
	"iec-60309-p-n-e-9h",
	"iec-60309-2p-e-4h",
	"iec-60309-2p-e-6h",
	"iec-60309-2p-e-9h",
	"iec-60309-3p-e-4h",
	"iec-60309-3p-e-6h",
	"iec-60309-3p-e-9h",
	"iec-60309-3p-n-e-4h",
	"iec-60309-3p-n-e-6h",
	"iec-60309-3p-n-e-9h",
	"iec-60906-1",
	"nbr-14136-10a",
	"nbr-14136-20a",
	"nema-1-15p",
	"nema-5-15p",
	"nema-5-20p",
	"nema-5-30p",
	"nema-5-50p",
	"nema-6-15p",
	"nema-6-20p",
	"nema-6-30p",
	"nema-6-50p",
	"nema-10-30p",
	"nema-10-50p",
	"nema-14-20p",
	"nema-14-30p",
	"nema-14-50p",
	"nema-14-60p",
	"nema-15-15p",
	"nema-15-20p",
	"nema-15-30p",
	"nema-15-50p",
	"nema-15-60p",
	"nema-l1-15p",
	"nema-l5-15p",
	"nema-l5-20p",
	"nema-l5-30p",
	"nema-l5-50p",
	"nema-l6-15p",
	"nema-l6-20p",
	"nema-l6-30p",
	"nema-l6-50p",
	"nema-l10-30p",
	"nema-l14-20p",
	"nema-l14-30p",
	"nema-l14-50p",
	"nema-l14-60p",
	"nema-l15-20p",
	"nema-l15-30p",
	"nema-l15-50p",
	"nema-l15-60p",
	"nema-l21-20p",
	"nema-l21-30p",
	"nema-l22-30p",
	"cs6361c",
	"cs6365c",
	"cs8165c",
	"cs8265c",
	"cs8365c",
	"cs8465c",
	"ita-c",
	"ita-e",
	"ita-f",
	"ita-ef",
	"ita-g",
	"ita-h",
	"ita-i",
	"ita-j",
	"ita-k",
	"ita-l",
	"ita-m",
	"ita-n",
	"ita-o",
	"usb-a",
	"usb-b",
	"usb-c",
	"usb-mini-a",
	"usb-mini-b",
	"usb-micro-a",
	"usb-micro-b",
	"usb-micro-ab",
	"usb-3-b",
	"usb-3-micro-b",
	"dc-terminal",
	"saf-d-grid",
	"neutrik-powercon-20",
	"neutrik-powercon-32",
	"neutrik-powercon-true1",
	"neutrik-powercon-true1-top",
	"ubiquiti-smartpower",
	"hardwired",
	"other",
	"",
}

func (v *PatchedWritablePowerPortTemplateRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PatchedWritablePowerPortTemplateRequestType(value)
	for _, existing := range AllowedPatchedWritablePowerPortTemplateRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PatchedWritablePowerPortTemplateRequestType", value)
}

// NewPatchedWritablePowerPortTemplateRequestTypeFromValue returns a pointer to a valid PatchedWritablePowerPortTemplateRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPatchedWritablePowerPortTemplateRequestTypeFromValue(v string) (*PatchedWritablePowerPortTemplateRequestType, error) {
	ev := PatchedWritablePowerPortTemplateRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PatchedWritablePowerPortTemplateRequestType: valid values are %v", v, AllowedPatchedWritablePowerPortTemplateRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PatchedWritablePowerPortTemplateRequestType) IsValid() bool {
	for _, existing := range AllowedPatchedWritablePowerPortTemplateRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PatchedWritablePowerPortTemplateRequest_type value
func (v PatchedWritablePowerPortTemplateRequestType) Ptr() *PatchedWritablePowerPortTemplateRequestType {
	return &v
}

type NullablePatchedWritablePowerPortTemplateRequestType struct {
	value *PatchedWritablePowerPortTemplateRequestType
	isSet bool
}

func (v NullablePatchedWritablePowerPortTemplateRequestType) Get() *PatchedWritablePowerPortTemplateRequestType {
	return v.value
}

func (v *NullablePatchedWritablePowerPortTemplateRequestType) Set(val *PatchedWritablePowerPortTemplateRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritablePowerPortTemplateRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritablePowerPortTemplateRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritablePowerPortTemplateRequestType(val *PatchedWritablePowerPortTemplateRequestType) *NullablePatchedWritablePowerPortTemplateRequestType {
	return &NullablePatchedWritablePowerPortTemplateRequestType{value: val, isSet: true}
}

func (v NullablePatchedWritablePowerPortTemplateRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritablePowerPortTemplateRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
