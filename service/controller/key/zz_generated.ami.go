// DO NOT EDIT. Generated with:
//
//    devctl@4.9.1-dev
//

package key

import "encoding/json"

var amiInfo = map[string]map[string]string{}

var amiJSON = []byte(`{
  "2191.5.0": {
    "ap-northeast-1": "ami-0abb8a024e2fca994",
    "ap-northeast-2": "ami-0da2d607f218e6330",
    "ap-south-1": "ami-08c1d433d7d880e2f",
    "ap-southeast-1": "ami-0ffc161cc11dae512",
    "ap-southeast-2": "ami-0e3812401bf255743",
    "ca-central-1": "ami-01a8fbed8281b5333",
    "eu-central-1": "ami-0f607d2a5a0fd8041",
    "eu-west-1": "ami-0c75aa1f3a9f34e91",
    "eu-west-2": "ami-0d21ee9c177f2562f",
    "eu-west-3": "ami-0a41df11d9c5b1fd8",
    "sa-east-1": "ami-0e2883834330d8b09",
    "us-east-1": "ami-0fba790bd1eea1719",
    "us-east-2": "ami-024b150f9a8d93151",
    "us-west-1": "ami-0871570295b52c558",
    "us-west-2": "ami-03617eef53cfa7a84"
  },
  "2247.5.0": {
    "ap-northeast-1": "ami-0df58a66c0b2fdc66",
    "ap-northeast-2": "ami-0ccb6ca76e1bd6092",
    "ap-south-1": "ami-0c0e023691b55dc61",
    "ap-southeast-1": "ami-061cbfd985a9d8f8e",
    "ap-southeast-2": "ami-083ffda3957a602fe",
    "ca-central-1": "ami-0c4c1c1e7cd65b2d2",
    "eu-central-1": "ami-0d4f8de8cdb783338",
    "eu-west-1": "ami-0feee01c888f0da46",
    "eu-west-2": "ami-01051a6731786b642",
    "eu-west-3": "ami-06c0b8dfeb7becf70",
    "sa-east-1": "ami-0bb50e7274175f3c2",
    "us-east-1": "ami-07ecdd174665b9180",
    "us-east-2": "ami-0a23274d2f7966f54",
    "us-west-1": "ami-015e4ecc50b60a5f5",
    "us-west-2": "ami-0cff3b30a036b4013"
  },
  "2247.6.0": {
    "ap-northeast-1": "ami-04313e11bc14cd1ff",
    "ap-northeast-2": "ami-0345f4d2e11635fa0",
    "ap-south-1": "ami-0f3feb88e63f87c87",
    "ap-southeast-1": "ami-0f17bcc3703db8419",
    "ap-southeast-2": "ami-0f102ff76cb29ff6d",
    "ca-central-1": "ami-07b90d921636638ec",
    "eu-central-1": "ami-0a8746e7c615c680c",
    "eu-west-1": "ami-044c1a858557a32e5",
    "eu-west-2": "ami-060ee621614d0cce1",
    "eu-west-3": "ami-0febe9d769fd5221f",
    "sa-east-1": "ami-015bfcdf1a7ff8f85",
    "us-east-1": "ami-0d2021fd8f91b4ef2",
    "us-east-2": "ami-045980983a94de38c",
    "us-west-1": "ami-04781ef85d776e35b",
    "us-west-2": "ami-0e11891fe6023ba51"
  },
  "2247.7.0": {
    "ap-northeast-1": "ami-03b739fc2d2e459a4",
    "ap-northeast-2": "ami-0a9276890527bce10",
    "ap-south-1": "ami-0fac33cc09f42345b",
    "ap-southeast-1": "ami-034dbfc3a57a9c5b7",
    "ap-southeast-2": "ami-052ff76e89468c8b7",
    "ca-central-1": "ami-06e8544033faf0627",
    "eu-central-1": "ami-0c89bf0ec0c4a179e",
    "eu-west-1": "ami-084ba4e37dbb4cbb3",
    "eu-west-2": "ami-087de5073a273f459",
    "eu-west-3": "ami-061080f901ff07c36",
    "sa-east-1": "ami-081aa7b8d0a257e72",
    "us-east-1": "ami-07a74c483029110e8",
    "us-east-2": "ami-00a0c9ddc32451f34",
    "us-west-1": "ami-0f9b55572bef76937",
    "us-west-2": "ami-0e374c8a8ae0cb51c"
  },
  "2303.3.0": {
    "ap-northeast-1": "ami-06f0f6d1803cebc00",
    "ap-northeast-2": "ami-04ca85723ec15146f",
    "ap-south-1": "ami-0fde66e0672da3c77",
    "ap-southeast-1": "ami-080daf9ba3a611f10",
    "ap-southeast-2": "ami-0814264bec66c0993",
    "ca-central-1": "ami-03cc5c255029bfab3",
    "eu-central-1": "ami-0d5a49659dec77965",
    "eu-west-1": "ami-0f86251689411749d",
    "eu-west-2": "ami-09071fd99e4d52356",
    "eu-west-3": "ami-0331176a89635da0c",
    "sa-east-1": "ami-061e80c52825921a4",
    "us-east-1": "ami-0bec30eaca6325366",
    "us-east-2": "ami-05596993ab3e27a8f",
    "us-west-1": "ami-01a92755bd145f1f0",
    "us-west-2": "ami-0448d339360a2373d"
  },
  "2303.3.1": {
    "ap-northeast-1": "ami-0394df49d3cb3b0f0",
    "ap-northeast-2": "ami-01ed6597adfa7baae",
    "ap-south-1": "ami-0330711a1827d99d1",
    "ap-southeast-1": "ami-0c4a6ebbb5dc59b73",
    "ap-southeast-2": "ami-0c3a27cfbf2a07830",
    "ca-central-1": "ami-009157164163bf01e",
    "eu-central-1": "ami-0215a3496f6a601d6",
    "eu-west-1": "ami-02815ba59fd1215da",
    "eu-west-2": "ami-0ccb9d395b9e03ff1",
    "eu-west-3": "ami-04dd061b93942c838",
    "sa-east-1": "ami-03441ef7cdc572d84",
    "us-east-1": "ami-0db7de4b628d58045",
    "us-east-2": "ami-0ddd9867c53c030da",
    "us-west-1": "ami-021144166741aaa1e",
    "us-west-2": "ami-014436c61e95d301f"
  },
  "2303.4.0": {
    "ap-northeast-1": "ami-041a65d4c661aa4fb",
    "ap-northeast-2": "ami-0e0ede38c27d8405b",
    "ap-south-1": "ami-0644110c2b1384834",
    "ap-southeast-1": "ami-0151c71bb035c8819",
    "ap-southeast-2": "ami-0abe98d4c655904d0",
    "ca-central-1": "ami-0a016f2796a1c8ccd",
    "eu-central-1": "ami-0cfe2a1aa7964bf28",
    "eu-west-1": "ami-06c2012b184b8dd26",
    "eu-west-2": "ami-08734a32635d94d0b",
    "eu-west-3": "ami-04e65366067632fd1",
    "sa-east-1": "ami-08a4cc66bbd2494d7",
    "us-east-1": "ami-0b040590c5b193e0a",
    "us-east-2": "ami-04de768ebf7c816b4",
    "us-west-1": "ami-01687a29be800de16",
    "us-west-2": "ami-0432fb48aa07baa3f"
  },
  "2345.3.0": {
    "ap-east-1": "ami-0a813620447e80b05",
    "ap-northeast-1": "ami-02af6d096f0a2f96c",
    "ap-northeast-2": "ami-0dd7a397031040ff1",
    "ap-south-1": "ami-03205de7c095444bb",
    "ap-southeast-1": "ami-0666b8c77a4148316",
    "ap-southeast-2": "ami-0be9b0ada4e9f0c7a",
    "ca-central-1": "ami-0b0044f3e521384ae",
    "eu-central-1": "ami-0c9ac894c7ec2e6dd",
    "eu-north-1": "ami-06420e2a1713889dd",
    "eu-west-1": "ami-09f0cd6af1e455cd9",
    "eu-west-2": "ami-0b1588137b7790e8c",
    "eu-west-3": "ami-01a8e028daf4d66cf",
    "me-south-1": "ami-0a6518241a90f491f",
    "sa-east-1": "ami-037e10c3bd117fb3e",
    "us-east-1": "ami-007776654941e2586",
    "us-east-2": "ami-0b0a4944bd30c6b85",
    "us-west-1": "ami-02fefca3d52d15b1d",
    "us-west-2": "ami-0390d41fd4e4a3529"
  },
  "2345.3.1": {
    "ap-east-1": "ami-0e28e38ecce552688",
    "ap-northeast-1": "ami-074891de68922e1f4",
    "ap-northeast-2": "ami-0a1a6a05c79bcdfe4",
    "ap-south-1": "ami-0765ae35424be8ad8",
    "ap-southeast-1": "ami-0f20e37280d5c8c5c",
    "ap-southeast-2": "ami-016e5e9a74cc6ef86",
    "ca-central-1": "ami-09afcf2e90761d6e6",
    "cn-north-1": "ami-019174dba14053d2a",
    "cn-northwest-1": "ami-004e81bc53b1e6ffa",
    "eu-central-1": "ami-0a9a5d2b65cce04eb",
    "eu-north-1": "ami-0bbfc19aa4c355fe2",
    "eu-west-1": "ami-002db020452770c0f",
    "eu-west-2": "ami-024928e37dcc18a42",
    "eu-west-3": "ami-083e4a190c9b050b1",
    "me-south-1": "ami-078eb26f287443167",
    "sa-east-1": "ami-01180d594d0315f65",
    "us-east-1": "ami-011655f166912d5ba",
    "us-east-2": "ami-0e30f3d8cbc900ff4",
    "us-west-1": "ami-0360d32ce24f1f05f",
    "us-west-2": "ami-0c1654a9988866a1f"
  },
  "2512.2.0": {
    "ap-east-1": "ami-0fc4cc3502335a06a",
    "ap-northeast-1": "ami-03b5d6612a5b2ec8b",
    "ap-northeast-2": "ami-0ec5815355e429952",
    "ap-south-1": "ami-08a33ecec11b45d1a",
    "ap-southeast-1": "ami-0922d45771d6fbc5c",
    "ap-southeast-2": "ami-0f03a2f9d16b9d2b9",
    "ca-central-1": "ami-00985d7f177e24c76",
    "cn-north-1": "ami-0c4e9c1d91f27ab3f",
    "cn-northwest-1": "ami-03ca6476c6483b09f",
    "eu-central-1": "ami-06071ac001b33824c",
    "eu-north-1": "ami-0bd13a88480e7b69e",
    "eu-west-1": "ami-06652549a5f44581a",
    "eu-west-2": "ami-0ba34e692d93d38d6",
    "eu-west-3": "ami-0bb5682ae45553922",
    "me-south-1": "ami-0c37b6687e8c30560",
    "sa-east-1": "ami-02c75f06066115138",
    "us-east-1": "ami-07e9d85a5777405b6",
    "us-east-2": "ami-095e6549e6d67c0cc",
    "us-west-1": "ami-055a1252b1382338e",
    "us-west-2": "ami-014b4b19cc0222265"
  },
  "2512.2.1": {
    "ap-east-1": "ami-0c6ce81b14dab9a88",
    "ap-northeast-1": "ami-095f21ed1ce4b1950",
    "ap-northeast-2": "ami-0d7eacf9d737bc5db",
    "ap-south-1": "ami-0af9f9860996f7d6b",
    "ap-southeast-1": "ami-0170beb17a0899c33",
    "ap-southeast-2": "ami-0e9e553ab0b54c8d2",
    "ca-central-1": "ami-0b529c1051737b2d9",
    "cn-north-1": "ami-0b9e98c5bc5c0bb3f",
    "cn-northwest-1": "ami-0c6b8814cc7d11a6f",
    "eu-central-1": "ami-0cefde98784480dd3",
    "eu-north-1": "ami-0e6370d6bd74415a9",
    "eu-west-1": "ami-08738afa3d25a1196",
    "eu-west-2": "ami-0c353732dda0c25c4",
    "eu-west-3": "ami-07416fa4faa74e021",
    "me-south-1": "ami-0762545e56fa8d992",
    "sa-east-1": "ami-0584c315c8d17a8a8",
    "us-east-1": "ami-01a83a3acf30b4638",
    "us-east-2": "ami-083c654d2f0469f9e",
    "us-west-1": "ami-03efdc564552f8ed5",
    "us-west-2": "ami-0bb54692374ac10a7"
  },
  "2512.3.0": {
    "ap-east-1": "ami-0b4716b871beafd11",
    "ap-northeast-1": "ami-04f213d4f66927487",
    "ap-northeast-2": "ami-0e5f1e55bab0bb760",
    "ap-south-1": "ami-0b2489dc089ea583d",
    "ap-southeast-1": "ami-0d8caf5936c905a21",
    "ap-southeast-2": "ami-0d53afed89a5f4b18",
    "ca-central-1": "ami-074872d5522133bfc",
    "cn-north-1": "ami-05824d6697a544711",
    "cn-northwest-1": "ami-04559bbba254b151e",
    "eu-central-1": "ami-0fe804e5404503f60",
    "eu-north-1": "ami-055e9b7a756995e3d",
    "eu-west-1": "ami-025de04fe63bc07f7",
    "eu-west-2": "ami-002b55b01c5d28271",
    "eu-west-3": "ami-0e7f5a1b98e008936",
    "me-south-1": "ami-0ea26ce6ee41986b2",
    "sa-east-1": "ami-09c0060c1f1df4309",
    "us-east-1": "ami-099272f49841e5958",
    "us-east-2": "ami-058b7413adffb72e6",
    "us-west-1": "ami-01e17a7bc153725e0",
    "us-west-2": "ami-0e9f87aba5d30d988"
  },
  "2512.4.0": {
    "ap-east-1": "ami-0d4b90cf21cd58b85",
    "ap-northeast-1": "ami-0b98da1ee6c19bc67",
    "ap-northeast-2": "ami-0a988bb9d61fb1826",
    "ap-south-1": "ami-03a5fe985c9851954",
    "ap-southeast-1": "ami-0d6c8aabed9f9c822",
    "ap-southeast-2": "ami-0a03d93cecf574ce1",
    "ca-central-1": "ami-0974f917c3751e5e7",
    "cn-north-1": "ami-0efeb4100200884e1",
    "cn-northwest-1": "ami-01decc4949b213fe7",
    "eu-central-1": "ami-01506ad207ee36e67",
    "eu-north-1": "ami-0fa92d3c761f246f4",
    "eu-west-1": "ami-0af21c037dd8874e9",
    "eu-west-2": "ami-0be5e79d719deeb60",
    "eu-west-3": "ami-0a907a1d3610e14e9",
    "me-south-1": "ami-0cedbefc8b1459527",
    "sa-east-1": "ami-0ba522327fc4b4b41",
    "us-east-1": "ami-094d167ed104bc0a2",
    "us-east-2": "ami-03f331e65835b00db",
    "us-west-1": "ami-0240d7c67d8558e56",
    "us-west-2": "ami-088a77186416491fd"
  },
  "2512.5.0": {
    "ap-east-1": "ami-03615c69897e4f3de",
    "ap-northeast-1": "ami-055871fdac0dd4e17",
    "ap-northeast-2": "ami-0b085a2ec12d8e571",
    "ap-south-1": "ami-0208aa7bdadb1d236",
    "ap-southeast-1": "ami-0dd6837c585679993",
    "ap-southeast-2": "ami-0199a2c6bec4989ec",
    "ca-central-1": "ami-09639bd23e29e2786",
    "cn-north-1": "ami-04a4240b897b2e912",
    "cn-northwest-1": "ami-0cacf0b65d427baa3",
    "eu-central-1": "ami-0c3a6e66ea411f9d3",
    "eu-north-1": "ami-0d1508e09130629da",
    "eu-west-1": "ami-02108d8af56ca407d",
    "eu-west-2": "ami-0e7468764badfd6dd",
    "eu-west-3": "ami-0eeb728f74c23463b",
    "me-south-1": "ami-0c606b7862ceb910e",
    "sa-east-1": "ami-0dd76fb7370b4cf83",
    "us-east-1": "ami-0b02550649e7b00a4",
    "us-east-2": "ami-0053dfa21a6632e4c",
    "us-west-1": "ami-0247b27396bb9347c",
    "us-west-2": "ami-0b75e2f157200889f"
  },
  "2605.10.0": {
    "ap-east-1": "ami-09994584fe01481ee",
    "ap-northeast-1": "ami-03953690ce04db418",
    "ap-northeast-2": "ami-0d87101e5caaa2cdf",
    "ap-south-1": "ami-084eda0f718377c91",
    "ap-southeast-1": "ami-0c7e61ab88b0ef372",
    "ap-southeast-2": "ami-0dd44194e0a0d4aa2",
    "ca-central-1": "ami-00fb77d6d44a70630",
    "cn-north-1": "ami-03bcb6c6b7f669e29",
    "cn-northwest-1": "ami-04b3644782642eb4d",
    "eu-central-1": "ami-0d7b1dc2bc79a8de0",
    "eu-north-1": "ami-010d1b95ab9639c99",
    "eu-west-1": "ami-06d0b66dfc2f16074",
    "eu-west-2": "ami-0854b454f8cd46262",
    "eu-west-3": "ami-0196fdb9e8fb96098",
    "me-south-1": "ami-02440dfe957fe1598",
    "sa-east-1": "ami-095a5db7d558052cc",
    "us-east-1": "ami-06c0ea08de59e22d6",
    "us-east-2": "ami-0c5036855d78983f2",
    "us-west-1": "ami-034f6920b6c1bf18a",
    "us-west-2": "ami-07ce1eff537ac6861"
  },
  "2605.11.0": {
    "ap-east-1": "ami-0a66670a332004614",
    "ap-northeast-1": "ami-0542e4dff8420d487",
    "ap-northeast-2": "ami-041e489f08aac9255",
    "ap-south-1": "ami-06b541ee787f9dcdf",
    "ap-southeast-1": "ami-0721703baa9b72cd8",
    "ap-southeast-2": "ami-0c797e0867680e1d5",
    "ca-central-1": "ami-005f63b0367244cbb",
    "cn-north-1": "ami-006e38ad40da9b657",
    "cn-northwest-1": "ami-04ca1ab74bc985ff2",
    "eu-central-1": "ami-083477fab62dc2eb0",
    "eu-north-1": "ami-0d8a20560dc4f6d77",
    "eu-west-1": "ami-0f0d86a1db55e0fdb",
    "eu-west-2": "ami-0105750d9768df0a9",
    "eu-west-3": "ami-0544e90383096dfe3",
    "me-south-1": "ami-0fb337227637bcf0d",
    "sa-east-1": "ami-0354636e64ec00881",
    "us-east-1": "ami-01d772a46fc16d4f0",
    "us-east-2": "ami-0e042550cf12fd3a7",
    "us-west-1": "ami-006da855103391bd8",
    "us-west-2": "ami-0bef8aecb982525a5"
  },
  "2605.12.0": {
    "ap-east-1": "ami-0e61d15ce4cd16733",
    "ap-northeast-1": "ami-06471d461a6d5985b",
    "ap-northeast-2": "ami-0a992ed1dbf215084",
    "ap-south-1": "ami-04d088e39b621a690",
    "ap-southeast-1": "ami-09a03da18297d02eb",
    "ap-southeast-2": "ami-0e315fc285b4d8cd6",
    "ca-central-1": "ami-091e748f539647cba",
    "cn-north-1": "ami-06099399505bf9947",
    "cn-northwest-1": "ami-0ccfc7d5fa8c0ca66",
    "eu-central-1": "ami-0cf901161e7881321",
    "eu-north-1": "ami-058e73e441974b258",
    "eu-west-1": "ami-06a2a39aecdc787d9",
    "eu-west-2": "ami-09bb42a11dc68f714",
    "eu-west-3": "ami-04b352739fc594629",
    "me-south-1": "ami-072f3b2c1b0799985",
    "sa-east-1": "ami-0efcd6d9aa8134569",
    "us-east-1": "ami-0bb4587614fc28248",
    "us-east-2": "ami-08b49e7d67e26d443",
    "us-west-1": "ami-0d36c772438c6f895",
    "us-west-2": "ami-0e844b4a872122e30"
  },
  "2605.5.0": {
    "ap-east-1": "ami-000be1df42017285f",
    "ap-northeast-1": "ami-0b83e79fd7638154e",
    "ap-northeast-2": "ami-084bc95aae5b82d3a",
    "ap-south-1": "ami-0878be91eefcd6b76",
    "ap-southeast-1": "ami-0240e92312205a70d",
    "ap-southeast-2": "ami-077bb3348ba9ceb5f",
    "ca-central-1": "ami-02f2b90b06ce8e310",
    "eu-central-1": "ami-0a4c856d96d3012a5",
    "eu-north-1": "ami-0b940dde54096dac8",
    "eu-west-1": "ami-0ef157ace1e313660",
    "eu-west-2": "ami-0d3168673d7448bbb",
    "eu-west-3": "ami-031a50d66c880e870",
    "me-south-1": "ami-07a4a1cf5913f7364",
    "sa-east-1": "ami-000fd50b6bb4162fe",
    "us-east-1": "ami-0b0b90473c097c55a",
    "us-east-2": "ami-0af8a4497a35daeab",
    "us-west-1": "ami-01cd2d1d7d6f8ba8f",
    "us-west-2": "ami-025490fa95dba89ff"
  },
  "2605.6.0": {
    "ap-east-1": "ami-00475a938138cb2dd",
    "ap-northeast-1": "ami-081eca51ab8bd4c05",
    "ap-northeast-2": "ami-018af97fa18ae9433",
    "ap-south-1": "ami-0fbed348e59f98c1d",
    "ap-southeast-1": "ami-0305f8c979b581e7a",
    "ap-southeast-2": "ami-0eec6e45836a77030",
    "ca-central-1": "ami-0285fcf3cb3a6c662",
    "cn-north-1": "ami-05f73f0f401725c9e",
    "cn-northwest-1": "ami-03928b27bd062c1e9",
    "eu-central-1": "ami-08bd2c646313e99a6",
    "eu-north-1": "ami-0f5b9e77b6a061f3b",
    "eu-west-1": "ami-0a7cbb93fdc72a8f2",
    "eu-west-2": "ami-0eb0281a48a8ff213",
    "eu-west-3": "ami-0d1fd1e23ed8ab34b",
    "me-south-1": "ami-047edcc6c1141c10b",
    "sa-east-1": "ami-0aa770aecb43e9394",
    "us-east-1": "ami-02e6719a20017c3b6",
    "us-east-2": "ami-0be67f8e881cc9c62",
    "us-west-1": "ami-068c98f123a9a1a0b",
    "us-west-2": "ami-0dcea309142a5bb25"
  },
  "2605.7.0": {
    "ap-east-1": "ami-01b2873b11b0a9eaf",
    "ap-northeast-1": "ami-0cd9601ac4cd932cd",
    "ap-northeast-2": "ami-0e32c3791de53888c",
    "ap-south-1": "ami-0bef7791c50296f26",
    "ap-southeast-1": "ami-0b42bf8a879da36b8",
    "ap-southeast-2": "ami-00ac857ac1a747f6c",
    "ca-central-1": "ami-05cfd0d8e3aec0ba9",
    "cn-north-1": "ami-04f9e6933f6fa15b9",
    "cn-northwest-1": "ami-0d111dd1105ac9280",
    "eu-central-1": "ami-02b69575cfea47081",
    "eu-north-1": "ami-0cb1343efebf1a811",
    "eu-west-1": "ami-05f2aa12f8e8997f1",
    "eu-west-2": "ami-0a08c9f8aa8f64e95",
    "eu-west-3": "ami-068ee5cc5f770656d",
    "me-south-1": "ami-07def786183737b23",
    "sa-east-1": "ami-0869ead2adfa90fd8",
    "us-east-1": "ami-0148952d51cec7bc2",
    "us-east-2": "ami-05f3f3f80ea032416",
    "us-west-1": "ami-0131be7c07d35d7ab",
    "us-west-2": "ami-0312c8ef72c898c8e"
  },
  "2605.8.0": {
    "ap-east-1": "ami-0a5e8db3f2d9a83d7",
    "ap-northeast-1": "ami-0da7ff5b0d88b0f70",
    "ap-northeast-2": "ami-001762e18e5b9ad65",
    "ap-south-1": "ami-045aea6d72d3198a9",
    "ap-southeast-1": "ami-04ea62937d9ce8e98",
    "ap-southeast-2": "ami-0934a29645bf8c9c3",
    "ca-central-1": "ami-03cf05aa25f3b8654",
    "cn-north-1": "ami-0cc1c2c81cd81a513",
    "cn-northwest-1": "ami-052021faebcd97de7",
    "eu-central-1": "ami-01b3ed9fb3861422d",
    "eu-north-1": "ami-0964e6edec73c754e",
    "eu-west-1": "ami-0ce8041e111a5e0c1",
    "eu-west-2": "ami-0de1cf6a241b43cb4",
    "eu-west-3": "ami-0cbf5357d24278be2",
    "me-south-1": "ami-0fd9927614cb049d9",
    "sa-east-1": "ami-02ac7cb1e91d1ecf4",
    "us-east-1": "ami-006c88b6ca69f7811",
    "us-east-2": "ami-058a15cede9a8dbe8",
    "us-west-1": "ami-0b553b7d6c8ae2e4d",
    "us-west-2": "ami-07ead16a3df6a7bd4"
  },
  "2605.9.0": {
    "ap-east-1": "ami-05e05fff5f03af8fd",
    "ap-northeast-1": "ami-0f4119d70697e9a6d",
    "ap-northeast-2": "ami-011b85517bd23f35d",
    "ap-south-1": "ami-0e6881b0135edaf99",
    "ap-southeast-1": "ami-08049d829c2c7bf93",
    "ap-southeast-2": "ami-060f5b26c4aaa1dc0",
    "ca-central-1": "ami-0310962a5303a5818",
    "cn-north-1": "ami-0db50e298b0facfba",
    "cn-northwest-1": "ami-0f826e7f769727165",
    "eu-central-1": "ami-0e3bca7c248d2ba67",
    "eu-north-1": "ami-0c80ccf1689ff6bc5",
    "eu-west-1": "ami-01005e70742a799fe",
    "eu-west-2": "ami-00942ee9583135032",
    "eu-west-3": "ami-04a36caaa7424a720",
    "me-south-1": "ami-0f5b9662c5b09d8c1",
    "sa-east-1": "ami-0f8a9432c2e773e50",
    "us-east-1": "ami-03a19681a5790093b",
    "us-east-2": "ami-0f414340882689d76",
    "us-west-1": "ami-0c3cb60ad349c1a8e",
    "us-west-2": "ami-0f1648473e0ade796"
  },
  "2765.2.0": {
    "ap-east-1": "ami-00f396e09218a790e",
    "ap-northeast-1": "ami-080b6ee40b299cd75",
    "ap-northeast-2": "ami-0ebf726fddd8e993d",
    "ap-south-1": "ami-0d20061115f9b1dc3",
    "ap-southeast-1": "ami-00990ae5d3300f4d6",
    "ap-southeast-2": "ami-0ad77b0b4e052128a",
    "ca-central-1": "ami-0742e3035eea60b9e",
    "cn-north-1": "ami-034a03e16c448b26d",
    "cn-northwest-1": "ami-004a81646b56af369",
    "eu-central-1": "ami-0ff4fa860a810db69",
    "eu-north-1": "ami-09ecb2cee014118d6",
    "eu-west-1": "ami-0632aea2bf90c1fb6",
    "eu-west-2": "ami-071a5358b970c9427",
    "eu-west-3": "ami-045c04cdd9d1d18ff",
    "me-south-1": "ami-0ae9d3e31cfdfffc4",
    "sa-east-1": "ami-0a80693c7ecc1e9ba",
    "us-east-1": "ami-0c4058071ecde7d81",
    "us-east-2": "ami-0064af5873560f276",
    "us-west-1": "ami-05033ebb665b29745",
    "us-west-2": "ami-034e40010ec01940c"
  },
  "2765.2.1": {
    "ap-east-1": "ami-077666b81fc0f2c67",
    "ap-northeast-1": "ami-0fb10d3775d1c8523",
    "ap-northeast-2": "ami-03968f60d20f538c3",
    "ap-south-1": "ami-0d4be6436592b34f0",
    "ap-southeast-1": "ami-083d8f8838190b67a",
    "ap-southeast-2": "ami-0e0451058c8cd91ae",
    "ca-central-1": "ami-0215ff3857e37f8c6",
    "cn-north-1": "ami-0bc96bac1fbeacdd0",
    "cn-northwest-1": "ami-0cb4041fb59e0fc7f",
    "eu-central-1": "ami-01ea2878f20901a1f",
    "eu-north-1": "ami-05ef1a68a28afb834",
    "eu-west-1": "ami-0a639a0efa71c38f8",
    "eu-west-2": "ami-04ba5ad5d5d5e0496",
    "eu-west-3": "ami-0bb0ec5df2dcbd7dc",
    "me-south-1": "ami-015c9ebf4057439b2",
    "sa-east-1": "ami-0102bc1c72fe5df51",
    "us-east-1": "ami-0b87b9d37fade4ee5",
    "us-east-2": "ami-0fe445ad962759b06",
    "us-west-1": "ami-0dc915a4e96651161",
    "us-west-2": "ami-0a7f2ffd5308cbfe7"
  },
  "2765.2.2": {
    "ap-east-1": "ami-08143d3c2b4edd3de",
    "ap-northeast-1": "ami-0795df04242f10569",
    "ap-northeast-2": "ami-0cd250e377f035e38",
    "ap-south-1": "ami-092ef496893d0b2c7",
    "ap-southeast-1": "ami-0b1bdba8af4f2e2d6",
    "ap-southeast-2": "ami-036913e1339e2b7d4",
    "ca-central-1": "ami-0e6df7ef58fcfaa5d",
    "cn-north-1": "ami-0362ea42049a5e40d",
    "cn-northwest-1": "ami-0947dd3d77312ab18",
    "eu-central-1": "ami-03c52ea5d6f9f7d37",
    "eu-north-1": "ami-0359c8b88220e81f2",
    "eu-west-1": "ami-023c3f828f1f53d52",
    "eu-west-2": "ami-0551d1417af39acc9",
    "eu-west-3": "ami-0c15b62ee506279a5",
    "me-south-1": "ami-0ecc8f9014468b96d",
    "sa-east-1": "ami-04879f2240402da1b",
    "us-east-1": "ami-0af56e57b9765d6b0",
    "us-east-2": "ami-09207b4b2c351fadc",
    "us-west-1": "ami-082b6b073b0d596b9",
    "us-west-2": "ami-02b46c73fed689d1c"
  },
  "2765.2.3": {
    "ap-east-1": "ami-0f3eb6124e1fef7ae",
    "ap-northeast-1": "ami-0ca22b80cac36c3db",
    "ap-northeast-2": "ami-05b94b1342f80e813",
    "ap-south-1": "ami-01fd9de6fdd16e972",
    "ap-southeast-1": "ami-08959a8939d1cb5a2",
    "ap-southeast-2": "ami-044295df63a278d0e",
    "ca-central-1": "ami-04bea8a9e8ddf35c1",
    "cn-north-1": "ami-051e3a69ffd349b50",
    "cn-northwest-1": "ami-0d9485d008b887bda",
    "eu-central-1": "ami-0b37b7aabd02e1586",
    "eu-north-1": "ami-0477be1fda3907a95",
    "eu-west-1": "ami-0dff9f0375c3cdbc4",
    "eu-west-2": "ami-0bd141149cc324600",
    "eu-west-3": "ami-03d36fcfe9b195856",
    "me-south-1": "ami-076b972cbea0f9871",
    "sa-east-1": "ami-0505708e3613d9872",
    "us-east-1": "ami-0eb8222969c04150e",
    "us-east-2": "ami-0fc996b00faa96d41",
    "us-west-1": "ami-0ae3c3cd0575045fb",
    "us-west-2": "ami-0654e052c52fe0398"
  },
  "2765.2.4": {
    "ap-east-1": "ami-001535916b3955e60",
    "ap-northeast-1": "ami-004476b3b178ee15d",
    "ap-northeast-2": "ami-0cdc7edbb441d1d0c",
    "ap-south-1": "ami-093d9c0c4ad4ad43e",
    "ap-southeast-1": "ami-0a7a0f34a9cb8e442",
    "ap-southeast-2": "ami-06966997b5abfa338",
    "ca-central-1": "ami-0fb2958af225a4510",
    "cn-north-1": "ami-00a260fcab94d7309",
    "cn-northwest-1": "ami-0c1ee11365e813241",
    "eu-central-1": "ami-02eae8cac7ba7d8cc",
    "eu-north-1": "ami-005677f8d981eadce",
    "eu-west-1": "ami-0110f218723a06667",
    "eu-west-2": "ami-03b95062b6d3c6b19",
    "eu-west-3": "ami-0f751c6f3c7d4a1e1",
    "me-south-1": "ami-024b0389de8f5e2a6",
    "sa-east-1": "ami-0fac99732f2a28fd2",
    "us-east-1": "ami-0c3c3f3925552a799",
    "us-east-2": "ami-03e4d01d25a4b90a1",
    "us-west-1": "ami-0fe9fa563274b9b74",
    "us-west-2": "ami-002e3f2c1bfc76d84"
  },
  "2765.2.5": {
    "ap-east-1": "ami-0d57e68826623475c",
    "ap-northeast-1": "ami-00198a70f3bcc88c3",
    "ap-northeast-2": "ami-06e3bf9845d7e5ae8",
    "ap-south-1": "ami-05cd2019a166d7667",
    "ap-southeast-1": "ami-0696cdbbe125cd6b5",
    "ap-southeast-2": "ami-05ba71789bc742de8",
    "ca-central-1": "ami-0950bbe112318bb72",
    "cn-north-1": "ami-0d8fbf01ac4f7c044",
    "cn-northwest-1": "ami-0f7b7db14a96fb472",
    "eu-central-1": "ami-0a4e8fe96096bd69e",
    "eu-north-1": "ami-0626a9a740c998534",
    "eu-west-1": "ami-0c2c969ff253ed355",
    "eu-west-2": "ami-0fddaf47f1bb8620b",
    "eu-west-3": "ami-0a708ddc6c1cd4804",
    "me-south-1": "ami-01c189e94f5a4a243",
    "sa-east-1": "ami-00dcabaeee211aa8d",
    "us-east-1": "ami-059f996f16c2906d6",
    "us-east-2": "ami-024dc2ad9683d233d",
    "us-west-1": "ami-0769e29fb2805c954",
    "us-west-2": "ami-08ba66bfba9b52612"
  },
  "2765.2.6": {
    "ap-east-1": "ami-0df249fad4e423ae7",
    "ap-northeast-1": "ami-028697e8df1ff071b",
    "ap-northeast-2": "ami-0c5b8f2d07d21da16",
    "ap-south-1": "ami-055b64c22dbcd61b0",
    "ap-southeast-1": "ami-045357ea038a43fe7",
    "ap-southeast-2": "ami-05df81d055054698d",
    "ca-central-1": "ami-0f639872bfcb49738",
    "cn-north-1": "ami-0a8f044a16ccf9b7f",
    "cn-northwest-1": "ami-0794ec943cd349815",
    "eu-central-1": "ami-055acc5a6e9587b44",
    "eu-north-1": "ami-04f64f11f4dacda92",
    "eu-west-1": "ami-019f09de46e4e3f88",
    "eu-west-2": "ami-0097d8b6241e9cf76",
    "eu-west-3": "ami-05b8b0131fbb39283",
    "me-south-1": "ami-0737c661a0881fd94",
    "sa-east-1": "ami-00bc3ae33287bc81b",
    "us-east-1": "ami-0fd66875fa1ef8395",
    "us-east-2": "ami-02eb704ee029f6b9e",
    "us-west-1": "ami-053fb35697f85574d",
    "us-west-2": "ami-019657181ea76e880"
  },
  "2905.2.0": {
    "ap-east-1": "ami-0ea76d2db0219a7cd",
    "ap-northeast-1": "ami-05bcd0fbf12fcb033",
    "ap-northeast-2": "ami-05ee37031c2748c06",
    "ap-south-1": "ami-0dd6cd77a73d99aa4",
    "ap-southeast-1": "ami-07653699aac271784",
    "ap-southeast-2": "ami-0bf3dfefc6bb42da9",
    "ca-central-1": "ami-0e62acc0aaf228bca",
    "cn-north-1": "ami-0b6a1ab0d4659b661",
    "cn-northwest-1": "ami-08eebe354e1b9ccee",
    "eu-central-1": "ami-0740fd6ee3280d143",
    "eu-north-1": "ami-02399139ae92419bf",
    "eu-west-1": "ami-00722d8c67ffe4d70",
    "eu-west-2": "ami-095b0a6b2be3dfbca",
    "eu-west-3": "ami-03cbd0fdb5e8ebef5",
    "me-south-1": "ami-014d38c867a7885c1",
    "sa-east-1": "ami-0bed9dc9392fa0459",
    "us-east-1": "ami-05f41d653298051ef",
    "us-east-2": "ami-0fa64dc57392b7d94",
    "us-west-1": "ami-0f7f23b07a38d8f40",
    "us-west-2": "ami-08b0dbd7ab8eee76a"
  },
  "2905.2.1": {
    "ap-east-1": "ami-0a9dcd7d0ec2a01d3",
    "ap-northeast-1": "ami-046935c916fafd3bf",
    "ap-northeast-2": "ami-0315db16dc51878ac",
    "ap-south-1": "ami-07fa212ec000724e7",
    "ap-southeast-1": "ami-02e775f8791c74491",
    "ap-southeast-2": "ami-03c592f02a51246ba",
    "ca-central-1": "ami-0d469dca314a604fc",
    "cn-north-1": "ami-0b6247c344ff61db3",
    "cn-northwest-1": "ami-06c9cd52476fbc496",
    "eu-central-1": "ami-0db0e33ee3b0f24e3",
    "eu-north-1": "ami-0460c67c852bc8bef",
    "eu-west-1": "ami-0f98595f12bb76504",
    "eu-west-2": "ami-0c2605ab7060f8536",
    "eu-west-3": "ami-0ec024c62c26da40f",
    "me-south-1": "ami-0a7bc7d16b687d1d8",
    "sa-east-1": "ami-0a846564314960047",
    "us-east-1": "ami-06c2716c74c43eedf",
    "us-east-2": "ami-031fa346b3be8a6a6",
    "us-west-1": "ami-0670e450bda89efc0",
    "us-west-2": "ami-0459c48186a4a3839"
  },
  "2905.2.2": {
    "ap-east-1": "ami-0fb69cb4856dc4ab6",
    "ap-northeast-1": "ami-09412388b5b424d27",
    "ap-northeast-2": "ami-0db3d1ed3fdcec712",
    "ap-south-1": "ami-0e75281de2170d4f4",
    "ap-southeast-1": "ami-082848295014c48d9",
    "ap-southeast-2": "ami-04e816ebd60602f7c",
    "ca-central-1": "ami-02ae346be0272c246",
    "cn-north-1": "ami-0e82aa97297980691",
    "cn-northwest-1": "ami-0f95df669661cc402",
    "eu-central-1": "ami-0a6d54eac69d76455",
    "eu-north-1": "ami-0a8c02e752f691f82",
    "eu-west-1": "ami-0d33c8bb79449ed26",
    "eu-west-2": "ami-0efbb2ab64a448d47",
    "eu-west-3": "ami-048d2ec98b7a9c558",
    "me-south-1": "ami-073d0b512070f4ed6",
    "sa-east-1": "ami-0d001bd415770f1ad",
    "us-east-1": "ami-03ed5cba16ab3bfbf",
    "us-east-2": "ami-0990ba8b5f406fd60",
    "us-west-1": "ami-0293bef8a5f6cf382",
    "us-west-2": "ami-028d0f76d5d168089"
  },
  "2905.2.3": {
    "ap-east-1": "ami-0970df5f3bb8c5a7f",
    "ap-northeast-1": "ami-0e80958ff6e9d796f",
    "ap-northeast-2": "ami-00ea81c55e7de2ed2",
    "ap-south-1": "ami-07e85bfb1780f3b22",
    "ap-southeast-1": "ami-02f9d13f3ca4da8b8",
    "ap-southeast-2": "ami-0d24ec44c7814e36c",
    "ca-central-1": "ami-0e2ff0c47370992b8",
    "cn-north-1": "ami-09029cafa902aec1f",
    "cn-northwest-1": "ami-03303b931671c8c32",
    "eu-central-1": "ami-0e49f094a7d43a596",
    "eu-north-1": "ami-006120476f26704b2",
    "eu-west-1": "ami-05388676a7f5dd170",
    "eu-west-2": "ami-0abfa0ac1f8d2394b",
    "eu-west-3": "ami-0c509e57c85cd9e94",
    "me-south-1": "ami-00af95fabb2edddb0",
    "sa-east-1": "ami-0bd00a4e87d2600b0",
    "us-east-1": "ami-0c0f11b04e5c90a62",
    "us-east-2": "ami-07e82385de8861b75",
    "us-west-1": "ami-01576391fa748ab56",
    "us-west-2": "ami-00ab35df371f04b39"
  }
}`)

func init() {
	err := json.Unmarshal(amiJSON, &amiInfo)
	if err != nil {
		panic(err)
	}
}
