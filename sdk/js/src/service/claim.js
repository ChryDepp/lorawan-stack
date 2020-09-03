// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import Marshaler from '../util/marshaler'

class DeviceClaim {
  constructor(api, { stackConfig }) {
    if (!api) {
      throw new Error('Cannot initialize DeviceClaim service without api object.')
    }
    this._api = api
    this._stackConfig = stackConfig
  }

  // Claim
  async claim(
    applicationId,
    qrCode,
    targetNetworkServerAddress = this._stackConfig.nsHost,
    targetApplicationServerAddress = this._stackConfig.asHost,
  ) {
    const payload = {
      qr_code: btoa(qrCode),
      target_application_ids: {
        application_id: applicationId,
      },
      target_network_server_address: targetNetworkServerAddress,
      target_application_server_address: targetApplicationServerAddress,
    }

    const response = await this._api.EndDeviceClaimingServer.Claim(undefined, payload)

    return Marshaler.payloadSingleResponse(response)
  }
}

export default DeviceClaim