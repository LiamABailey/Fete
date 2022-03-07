# Fête: A Lightweight Feature Definition Store

## About Fête
Fête provides a simple, standardized interface to store and serve feature *definitions*. While many feature store solutions rely on a repeated call and response to support deployed models, Fête enables the deployment of feature *definitions* to the model pre-compute resource.

![Fête vs Feature Stores](https://github.com/LiamABailey/Fete/blob/main/readme/f%C3%AAte_usage.png?raw=true)

As can be seen above, the core usage pattern in Fête is not dissimilar to many other traditional feature stores. The key difference comes in production: instead of frequently requesting feature values from a feature store, with Fête feature *definitions* are only requested at pipeline setup, helping to limit network traffic.

## Features
- **Feature Definition Storage** : Interact with instructions, not values
- **Versioning** : Simultaneously maintain multiple versions of a transformation
- **Tagging** : Tag feature definitions to support exploration and documentation
- **Testing** : Support for automated test execution at time-of-request
- **Auditability** : Logs *which* definition was requested *when* by *who*

## Should I Use Fête for my project?
Fête, as an alternative to traditional feature stores, offers the following benefits and drawbacks. Depending on the specific use case, Fête may not be the best tool for the job.
### Key Differentiators:
- With Fête, feature definitions only need be requested once per model pipeline instance: at the build step. For traditional feature stores, data is constantly sent to and from the service.
- Because Fête doesn't operate on raw data, you limit transmission volumes and maintain greater control of your data.
- Traditional feature stores provide out-of-the-gate separation of feature computation and prediction computation. Although Fête can be wrapped by a computational layer, it is biased towards feature computation alongside prediction computation.
- Fête doesn't operate directly on raw data, and consequently it's agnostic to the transformation type (batch vs streaming vs on-demand).
- Fête does not serve as a route to data quality monitoring (some feature store services offer monitoring capabilities).
