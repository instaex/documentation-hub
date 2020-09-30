﻿using System;

namespace BlockM3.AEternity.SDK.Sophia.Attributes
{
    [AttributeUsage(AttributeTargets.Class | AttributeTargets.Property | AttributeTargets.Parameter | AttributeTargets.GenericParameter)]
    public class SophiaStateAttribute : Attribute, ISophiaAttribute
    {
    }
}